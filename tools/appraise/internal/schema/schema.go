// Package schema wires the appraise CLI domain onto an agentquery.Schema.
//
// The appraise tool is unusual: its operations (calc, list, summary) don't read
// from a homogeneous item list. Instead, each operation does its own work
// (running a calculator, listing modules, etc.). The Schema[CalcResult] type
// parameter is used for field validation and projection, but the loader is not
// the primary data source — operations are self-contained.
package schema

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ivalx1s/skill-agent-facing-api/agentquery"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// CalcResult is the Schema type parameter. It holds the envelope fields
// that agentquery uses for field validation and projection. The actual
// calculation results are heterogeneous and projected via JSON roundtrip.
type CalcResult struct {
	Module         string
	Function       string
	Value          any
	Details        map[string]any
	Error          *string
	Interpretation string
}

// New creates and configures the agentquery Schema for the appraise CLI.
// dataDir is the root directory for grep/search functionality.
func New(dataDir string) *agentquery.Schema[CalcResult] {
	schema := agentquery.NewSchema[CalcResult](
		agentquery.WithDataDir(dataDir),
		agentquery.WithExtensions(".json", ".md"),
	)

	registerFields(schema)
	registerPresets(schema)
	schema.DefaultFields("default")

	// No loader needed — operations handle their own data.
	schema.SetLoader(func() ([]CalcResult, error) {
		return nil, nil
	})

	registry := calculators.NewRegistry()
	registerOperations(schema, registry)

	return schema
}

// registerFields registers all valid field names for parse-time validation.
// Accessors extract from CalcResult for envelope fields; detailed calc fields
// are only used for parse validation (projection is done via JSON roundtrip).
func registerFields(schema *agentquery.Schema[CalcResult]) {
	// --- Envelope fields (used for both validation and projection) ---
	schema.Field("module", func(r CalcResult) any { return r.Module })
	schema.Field("function", func(r CalcResult) any { return r.Function })
	schema.Field("value", func(r CalcResult) any { return r.Value })
	schema.Field("details", func(r CalcResult) any { return r.Details })
	schema.Field("error", func(r CalcResult) any { return r.Error })
	schema.Field("interpretation", func(r CalcResult) any { return r.Interpretation })

	// --- BVR fields ---
	noop := func(CalcResult) any { return nil }
	for _, f := range []string{
		"bvr", "standalone_sum", "bundle_price", "component_values",
	} {
		schema.Field(f, noop)
	}

	// --- Tier gap fields ---
	for _, f := range []string{
		"gaps", "from_tier", "to_tier", "price_gap_abs", "price_gap_pct",
		"value_gap", "value_to_price", "diagnosis",
	} {
		schema.Field(f, noop)
	}

	// --- LFK fields ---
	for _, f := range []string{
		"classifications", "leaders_count", "fillers_count", "killers_count",
		"classification", "rationale",
	} {
		schema.Field(f, noop)
	}

	// --- Dead weight fields ---
	for _, f := range []string{
		"dead_weight_ratio", "threshold", "passes", "dead_weight", "component_usage",
	} {
		schema.Field(f, noop)
	}

	// --- Financial fields ---
	for _, f := range []string{
		"revenue_per_customer", "cost_per_customer", "margin_per_customer",
		"margin_pct", "viable", "clv", "break_even_units", "fixed_costs",
		"contribution_margin",
	} {
		schema.Field(f, noop)
	}

	// --- Stress test fields ---
	for _, f := range []string{
		"base_margin", "stressed_margin", "survives_stress",
	} {
		schema.Field(f, noop)
	}

	// --- Cannibalization fields ---
	for _, f := range []string{
		"net_revenue_delta", "net_positive",
	} {
		schema.Field(f, noop)
	}

	// --- Scoring fields ---
	for _, f := range []string{
		"weighted_score", "decision", "dimensions", "weights_used",
		"risks", "avg_score", "max_score", "high_risks",
	} {
		schema.Field(f, noop)
	}

	// --- Schema/introspection fields ---
	for _, f := range []string{
		"modules", "functions", "schema",
	} {
		schema.Field(f, noop)
	}
}

// registerPresets sets up the named field bundles.
func registerPresets(schema *agentquery.Schema[CalcResult]) {
	schema.Preset("minimal", "value", "interpretation")
	schema.Preset("default", "module", "function", "value", "interpretation")
	schema.Preset("full", "module", "function", "value", "details", "interpretation", "error")
}

// registerOperations registers all DSL operations on the schema.
func registerOperations(schema *agentquery.Schema[CalcResult], registry *calculators.Registry) {

	// --- calc ---
	schema.OperationWithMetadata("calc", func(ctx agentquery.OperationContext[CalcResult]) (any, error) {
		return execCalc(ctx, registry)
	}, agentquery.OperationMetadata{
		Description: "Run a calculation by module.function",
		Parameters: []agentquery.ParameterDef{
			{Name: "module.function", Type: "string", Optional: false, Description: "Calculator to run (e.g. pricing.bvr)"},
			{Name: "input", Type: "string", Optional: true, Description: "Path to input JSON file"},
		},
		Examples: []string{
			`calc(pricing.bvr, input="data.json")`,
			`calc(pricing.bvr, input="data.json") { minimal }`,
			`calc(financial.clv, input="data.json") { value }`,
		},
	})

	// --- list ---
	schema.OperationWithMetadata("list", func(ctx agentquery.OperationContext[CalcResult]) (any, error) {
		return execList(ctx, registry)
	}, agentquery.OperationMetadata{
		Description: "List available modules or functions",
		Parameters: []agentquery.ParameterDef{
			{Name: "what", Type: "string", Optional: false, Description: "'modules' or 'functions' (positional)"},
			{Name: "module", Type: "string", Optional: true, Description: "Filter functions by module name"},
		},
		Examples: []string{
			"list(modules)",
			"list(functions, module=pricing)",
			"list(functions)",
		},
	})

	// --- summary ---
	schema.OperationWithMetadata("summary", func(ctx agentquery.OperationContext[CalcResult]) (any, error) {
		return execSummary(registry)
	}, agentquery.OperationMetadata{
		Description: "Module/function counts overview",
		Examples: []string{
			"summary()",
		},
	})
}

// --- Operation implementations ---

// execCalc handles: calc(module.function, input="/path.json") { fields }
func execCalc(ctx agentquery.OperationContext[CalcResult], registry *calculators.Registry) (any, error) {
	if len(ctx.Statement.Args) == 0 {
		return nil, &agentquery.Error{
			Code:    agentquery.ErrValidation,
			Message: "calc requires at least one argument: module.function",
		}
	}

	// First positional arg is module.function
	moduleFunc := ctx.Statement.Args[0].Value
	module, function, err := parseModuleFunction(moduleFunc)
	if err != nil {
		return nil, &agentquery.Error{
			Code:    agentquery.ErrValidation,
			Message: err.Error(),
		}
	}

	// Look for input= arg
	inputFile := ""
	for _, arg := range ctx.Statement.Args[1:] {
		if arg.Key == "input" {
			inputFile = arg.Value
		}
	}

	// Load input data
	input, err := loadInput(inputFile)
	if err != nil {
		return nil, &agentquery.Error{
			Code:    agentquery.ErrInternal,
			Message: fmt.Sprintf("loading input: %v", err),
		}
	}

	// Execute calculation
	result, err := registry.Execute(module, function, input)
	if err != nil {
		return nil, &agentquery.Error{
			Code:    agentquery.ErrInternal,
			Message: err.Error(),
		}
	}

	// Apply field projection if requested
	if len(ctx.Statement.Fields) > 0 {
		return projectResult(ctx.Selector, result)
	}

	return result, nil
}

// execList handles: list(modules) or list(functions, module=pricing)
func execList(ctx agentquery.OperationContext[CalcResult], registry *calculators.Registry) (any, error) {
	if len(ctx.Statement.Args) == 0 {
		return nil, &agentquery.Error{
			Code:    agentquery.ErrValidation,
			Message: "list requires argument: 'modules' or 'functions'",
		}
	}

	what := ctx.Statement.Args[0].Value
	switch what {
	case "modules":
		return map[string]any{
			"modules": registry.Modules(),
		}, nil
	case "functions":
		module := ""
		for _, arg := range ctx.Statement.Args {
			if arg.Key == "module" {
				module = arg.Value
			}
		}
		if module == "" {
			// Return all functions for all modules
			result := make(map[string]any)
			for _, m := range registry.Modules() {
				fns, _ := registry.Functions(m)
				result[m] = fns
			}
			return result, nil
		}
		fns, err := registry.Functions(module)
		if err != nil {
			return nil, &agentquery.Error{
				Code:    agentquery.ErrNotFound,
				Message: err.Error(),
			}
		}
		return map[string]any{
			"module":    module,
			"functions": fns,
		}, nil
	default:
		return nil, &agentquery.Error{
			Code:    agentquery.ErrValidation,
			Message: fmt.Sprintf("list argument must be 'modules' or 'functions', got %q", what),
		}
	}
}

// execSummary handles: summary() — module count + function count overview
func execSummary(registry *calculators.Registry) (any, error) {
	modules := registry.Modules()
	totalFuncs := 0
	moduleSummary := make(map[string]int)
	for _, m := range modules {
		fns, _ := registry.Functions(m)
		moduleSummary[m] = len(fns)
		totalFuncs += len(fns)
	}
	return map[string]any{
		"modules":         len(modules),
		"total_functions": totalFuncs,
		"per_module":      moduleSummary,
	}, nil
}

// --- Helpers ---

// parseModuleFunction splits "pricing.bvr" into ("pricing", "bvr").
func parseModuleFunction(s string) (string, string, error) {
	for i, ch := range s {
		if ch == '.' {
			if i == 0 || i == len(s)-1 {
				return "", "", fmt.Errorf("invalid module.function: %q", s)
			}
			return s[:i], s[i+1:], nil
		}
	}
	return "", "", fmt.Errorf("module.function format required (e.g. pricing.bvr), got %q", s)
}

// loadInput reads and parses an AppraisalInput from a JSON file.
// If path is empty, reads from stdin.
func loadInput(path string) (*domain.AppraisalInput, error) {
	var data []byte
	var err error

	if path == "" || path == "-" {
		data, err = readStdin()
	} else {
		data, err = os.ReadFile(path)
	}
	if err != nil {
		return nil, err
	}

	var input domain.AppraisalInput
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, fmt.Errorf("parsing input JSON: %w", err)
	}
	return &input, nil
}

// readStdin reads all of stdin. Returns error if stdin is a terminal (no pipe).
func readStdin() ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, fmt.Errorf("no input file specified and stdin is a terminal; use --input or pipe JSON")
	}

	var buf []byte
	tmp := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
		}
		if err != nil {
			break
		}
	}
	return buf, nil
}

// projectResult applies field projection to a heterogeneous calc result
// via JSON roundtrip. This handles the fact that different calculators
// return different result types (BVRResult, TierGapResult, etc.).
func projectResult(selector *agentquery.FieldSelector[CalcResult], result any) (map[string]any, error) {
	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		// Result might be an array or primitive — wrap it
		return map[string]any{"value": result}, nil
	}

	// Check if "full" preset is in effect — all fields selected means pass-through.
	// We can detect this by checking if there are more fields selected than
	// exist in the result map.
	selectedFields := selector.Fields()

	filtered := make(map[string]any)
	for k, v := range m {
		for _, f := range selectedFields {
			if f == k {
				filtered[k] = v
				break
			}
		}
	}

	return filtered, nil
}

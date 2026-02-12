// Package query implements DSL operation handlers.
package query

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/fields"
)

// Executor runs parsed DSL queries against the calculator registry.
type Executor struct {
	registry *calculators.Registry
}

// NewExecutor creates a new query executor.
func NewExecutor() *Executor {
	return &Executor{
		registry: calculators.NewRegistry(),
	}
}

// ExecuteStatement handles a single parsed statement.
func (e *Executor) ExecuteStatement(stmt Statement) (interface{}, error) {
	switch stmt.Operation {
	case "calc":
		return e.execCalc(stmt)
	case "list":
		return e.execList(stmt)
	case "schema":
		return e.execSchema(stmt)
	case "summary":
		return e.execSummary(stmt)
	default:
		return nil, fmt.Errorf("unknown operation: %s", stmt.Operation)
	}
}

// execCalc handles: calc(module.function, input="/path/to/data.json") { fields }
func (e *Executor) execCalc(stmt Statement) (interface{}, error) {
	if len(stmt.Args) == 0 {
		return nil, fmt.Errorf("calc requires at least one argument: module.function")
	}

	// First positional arg is module.function
	moduleFunc := stmt.Args[0].Value
	module, function, err := parseModuleFunction(moduleFunc)
	if err != nil {
		return nil, err
	}

	// Look for input= arg
	inputFile := ""
	for _, arg := range stmt.Args[1:] {
		if arg.Key == "input" {
			inputFile = arg.Value
		}
	}

	// Load input data
	input, err := loadInput(inputFile)
	if err != nil {
		return nil, fmt.Errorf("loading input: %w", err)
	}

	// Execute calculation
	result, err := e.registry.Execute(module, function, input)
	if err != nil {
		return nil, err
	}

	// Apply field projection if requested
	if len(stmt.Fields) > 0 {
		selector, err := fields.NewSelector(stmt.Fields)
		if err != nil {
			return nil, err
		}
		return fields.Project(selector, result)
	}

	return result, nil
}

// execList handles: list(modules) or list(functions, module=pricing)
func (e *Executor) execList(stmt Statement) (interface{}, error) {
	if len(stmt.Args) == 0 {
		return nil, fmt.Errorf("list requires argument: 'modules' or 'functions'")
	}

	what := stmt.Args[0].Value
	switch what {
	case "modules":
		return map[string]interface{}{
			"modules": e.registry.Modules(),
		}, nil
	case "functions":
		module := ""
		for _, arg := range stmt.Args {
			if arg.Key == "module" {
				module = arg.Value
			}
		}
		if module == "" {
			// Return all functions for all modules
			result := make(map[string]interface{})
			for _, m := range e.registry.Modules() {
				fns, _ := e.registry.Functions(m)
				result[m] = fns
			}
			return result, nil
		}
		fns, err := e.registry.Functions(module)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"module":    module,
			"functions": fns,
		}, nil
	default:
		return nil, fmt.Errorf("list argument must be 'modules' or 'functions', got %q", what)
	}
}

// execSchema handles: schema() or schema(module)
func (e *Executor) execSchema(stmt Statement) (interface{}, error) {
	result := make(map[string]interface{})
	for _, m := range e.registry.Modules() {
		fns, _ := e.registry.Functions(m)
		result[m] = fns
	}
	return map[string]interface{}{
		"schema": result,
	}, nil
}

// execSummary handles: summary() — module count + function count overview
func (e *Executor) execSummary(stmt Statement) (interface{}, error) {
	modules := e.registry.Modules()
	totalFuncs := 0
	moduleSummary := make(map[string]int)
	for _, m := range modules {
		fns, _ := e.registry.Functions(m)
		moduleSummary[m] = len(fns)
		totalFuncs += len(fns)
	}
	return map[string]interface{}{
		"modules":         len(modules),
		"total_functions": totalFuncs,
		"per_module":      moduleSummary,
	}, nil
}

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

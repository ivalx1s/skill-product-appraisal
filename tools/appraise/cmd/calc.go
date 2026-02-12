package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
	"github.com/spf13/cobra"
)

var (
	calcInputFile string
	calcFormat    string
)

var calcCmd = &cobra.Command{
	Use:   "calc <module> <function>",
	Short: "Run a calculation directly",
	Long: `Execute a single calculation by specifying module and function names.

Modules: pricing, bundle, financial, customer, product, scoring

Examples:
  appraise calc pricing bvr --input data.json
  appraise calc financial clv --input data.json
  appraise calc scoring go_no_go --input scoring.json
  appraise calc pricing bvr --input data.json --format compact

Use 'appraise q "list(modules)"' or 'appraise q "list(functions, module=pricing)"'
to discover available modules and functions.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		module := args[0]
		function := args[1]

		// Load input
		input, err := loadInputFile(calcInputFile)
		if err != nil {
			return fmt.Errorf("loading input: %w", err)
		}

		// Execute
		registry := calculators.NewRegistry()
		result, err := registry.Execute(module, function, input)
		if err != nil {
			return fmt.Errorf("calculation failed: %w", err)
		}

		// Output
		if calcFormat == "compact" || calcFormat == "llm" {
			return outputCalcCompact(result)
		}
		return outputCalcJSON(result)
	},
}

func init() {
	calcCmd.Flags().StringVar(&calcInputFile, "input", "", "Path to input JSON file (required)")
	calcCmd.Flags().StringVar(&calcFormat, "format", "json", "Output format: json (default) | compact | llm")
	_ = calcCmd.MarkFlagRequired("input")
}

func loadInputFile(path string) (*domain.AppraisalInput, error) {
	var data []byte
	var err error

	if path == "" || path == "-" {
		// Read from stdin
		info, statErr := os.Stdin.Stat()
		if statErr != nil {
			return nil, statErr
		}
		if info.Mode()&os.ModeCharDevice != 0 {
			return nil, fmt.Errorf("no input file specified and stdin is a terminal")
		}
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("reading stdin: %w", err)
		}
	} else {
		data, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
	}

	var input domain.AppraisalInput
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, fmt.Errorf("parsing input JSON: %w", err)
	}
	return &input, nil
}

func outputCalcJSON(result interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(result)
}

func outputCalcCompact(result interface{}) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err == nil {
		printCompactMap("", m)
		return nil
	}

	fmt.Fprintln(os.Stdout, string(data))
	return nil
}

func printCompactMap(prefix string, m map[string]interface{}) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch tv := v.(type) {
		case map[string]interface{}:
			printCompactMap(key, tv)
		case []interface{}:
			for i, item := range tv {
				itemKey := fmt.Sprintf("%s[%d]", key, i)
				if nested, ok := item.(map[string]interface{}); ok {
					printCompactMap(itemKey, nested)
				} else {
					fmt.Fprintf(os.Stdout, "%s:%v\n", itemKey, compactValue(item))
				}
			}
		default:
			fmt.Fprintf(os.Stdout, "%s:%v\n", key, compactValue(v))
		}
	}
}

func compactValue(v interface{}) string {
	switch tv := v.(type) {
	case string:
		return tv
	case float64:
		if tv == float64(int64(tv)) {
			return fmt.Sprintf("%d", int64(tv))
		}
		return fmt.Sprintf("%.4f", tv)
	case bool:
		if tv {
			return "true"
		}
		return "false"
	case nil:
		return ""
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

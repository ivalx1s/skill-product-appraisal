package cmd

import (
	"fmt"
	"os"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/query"
	"github.com/spf13/cobra"
)

var (
	queryFormat string
)

var queryCmd = &cobra.Command{
	Use:   "q '<query>'",
	Short: "Execute DSL queries against calculator modules",
	Long: `Execute structured DSL queries for calculations and schema introspection.

Syntax: operation(args) { fields }
Batching: semicolons separate multiple queries in one call.

Operations:
  calc(module.function, input="/path.json") { fields }   Run calculation
  list(modules)                                           List available modules
  list(functions, module=pricing)                         List functions in module
  schema()                                                Full schema overview
  summary()                                               Module/function counts

Field projection:
  { value interpretation }     Specific fields
  { minimal }                  Preset: value + interpretation
  { default }                  Preset: module + function + value + interpretation
  { full }                     Preset: all fields

Examples:
  appraise q 'calc(pricing.bvr, input="data.json")'
  appraise q 'list(modules)'
  appraise q 'calc(pricing.bvr, input="d.json") { value }; calc(pricing.bundle_discount, input="d.json") { value }'
  appraise q 'summary()'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := query.ModeJSON
		if queryFormat == "compact" || queryFormat == "llm" {
			mode = query.ModeCompact
		}

		err := query.ExecuteBatch(args[0], mode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "query failed: %v\n", err)
			os.Exit(1)
		}
		return nil
	},
}

func init() {
	queryCmd.Flags().StringVar(&queryFormat, "format", "json", "Output format: json (default) | compact | llm")
}

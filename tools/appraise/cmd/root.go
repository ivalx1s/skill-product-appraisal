// Package cmd implements the cobra CLI commands for the appraise tool.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ivalx1s/skill-agent-facing-api/agentquery"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/schema"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "appraise",
	Short: "Product and bundle appraisal calculator",
	Long: `appraise — agent-facing CLI for product/bundle evaluation calculations.

Three access modes:
  appraise q '<query>'                          DSL queries (structured reads)
  appraise grep '<pattern>'                     Scoped text search
  appraise calc <module> <function> --input <f> Direct calculation`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Data directory for grep search — default to current directory.
	dataDir := "."

	// Build schema with agentquery.
	s := schema.New(dataDir)

	// Custom q command: normalizes dotted identifiers + --format defaults to "json".
	rootCmd.AddCommand(queryCommand(s))

	// Custom grep command: matches old CLI interface (text output by default).
	rootCmd.AddCommand(grepCommand(s))

	// Keep the direct calc command.
	rootCmd.AddCommand(calcCmd)
}

// queryCommand creates a "q" subcommand with query normalization.
// The old parser allowed dots in identifiers (e.g. calc(pricing.bvr)).
// agentquery's parser does not, so we normalize dotted args to quoted strings.
// --format defaults to "json" (matches old behavior).
func queryCommand(s *agentquery.Schema[schema.CalcResult]) *cobra.Command {
	var format string

	cmd := &cobra.Command{
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
  appraise q 'calc(pricing.bvr, input="data.json")' --format json
  appraise q 'list(modules)' --format json
  appraise q 'calc(pricing.bvr, input="d.json") { value }; summary()' --format json
  appraise q 'summary()' --format compact`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			mode, err := parseOutputMode(format)
			if err != nil {
				return err
			}

			// Normalize: quote dotted identifiers for agentquery compatibility.
			normalized := schema.NormalizeQuery(args[0])

			data, err := s.QueryJSONWithMode(normalized, mode)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return err
		},
	}

	cmd.Flags().StringVar(&format, "format", "json", "Output format: json (default) | compact | llm")
	return cmd
}

// grepCommand creates a "grep" subcommand that searches data files.
// Matches the old CLI interface: text output by default, optional --format flag.
func grepCommand(s *agentquery.Schema[schema.CalcResult]) *cobra.Command {
	var (
		fileGlob        string
		caseInsensitive bool
		contextLines    int
		format          string
		dir             string
	)

	cmd := &cobra.Command{
		Use:   "grep '<pattern>'",
		Short: "Scoped full-text search across data files",
		Long: `Search for regex patterns in JSON and markdown data files.

Output format: path:line:content (ripgrep-style) by default.

Examples:
  appraise grep "bundle_value"
  appraise grep "bvr" --file "*.json"
  appraise grep "premium" -i -C 2
  appraise grep "cost" --dir ./data --file "*.json"
  appraise grep "cost" --format json`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := agentquery.SearchOptions{
				FileGlob:        fileGlob,
				CaseInsensitive: caseInsensitive,
				ContextLines:    contextLines,
			}

			// If --dir is specified, create a temporary schema with that dir.
			searchSchema := s
			if dir != "" {
				searchSchema = schema.New(dir)
			}

			// Determine output mode
			if format == "json" || format == "compact" || format == "llm" {
				mode, err := parseOutputMode(format)
				if err != nil {
					return err
				}
				data, err := searchSchema.SearchJSONWithMode(args[0], opts, mode)
				if err != nil {
					return err
				}
				_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
				return err
			}

			// Default: text output (matches old behavior)
			results, err := searchSchema.Search(args[0], opts)
			if err != nil {
				return fmt.Errorf("grep failed: %w", err)
			}

			if len(results) == 0 {
				fmt.Fprintln(os.Stderr, "no matches found")
				return nil
			}

			// Print in ripgrep-style text format
			prevPath := ""
			for _, r := range results {
				if !r.IsMatch {
					continue
				}
				if r.Source.Path != prevPath {
					if prevPath != "" {
						fmt.Println()
					}
					prevPath = r.Source.Path
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%s:%d:%s\n", r.Source.Path, r.Source.Line, r.Content)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&fileGlob, "file", "", "File name glob filter (e.g. '*.json')")
	cmd.Flags().BoolVarP(&caseInsensitive, "ignore-case", "i", false, "Case-insensitive search")
	cmd.Flags().IntVarP(&contextLines, "context", "C", 0, "Context lines before and after match")
	cmd.Flags().StringVar(&dir, "dir", "", "Search directory (default: current directory)")
	cmd.Flags().StringVar(&format, "format", "", "Output format: text (default) | json | compact | llm")
	return cmd
}

// parseOutputMode converts a string flag value to an OutputMode.
func parseOutputMode(s string) (agentquery.OutputMode, error) {
	switch strings.ToLower(s) {
	case "compact", "llm":
		return agentquery.LLMReadable, nil
	case "json":
		return agentquery.HumanReadable, nil
	default:
		return 0, fmt.Errorf("unknown format %q: use \"json\", \"compact\", or \"llm\"", s)
	}
}

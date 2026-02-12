package cmd

import (
	"fmt"
	"os"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/search"
	"github.com/spf13/cobra"
)

var (
	grepFile           string
	grepCaseInsensitive bool
	grepContext         int
	grepDir             string
)

var grepCmd = &cobra.Command{
	Use:   "grep '<pattern>'",
	Short: "Scoped full-text search across data files",
	Long: `Search for regex patterns in JSON and markdown data files.

Output format: path:line:content (ripgrep-style)

Examples:
  appraise grep "bundle_value"
  appraise grep "bvr" --file "*.json"
  appraise grep "premium" -i -C 2
  appraise grep "cost" --dir ./data --file "*.json"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := grepDir
		if dir == "" {
			dir = "."
		}

		opts := search.Options{
			FileGlob:        grepFile,
			CaseInsensitive: grepCaseInsensitive,
			ContextLines:    grepContext,
		}

		matches, err := search.Grep(dir, args[0], opts)
		if err != nil {
			return fmt.Errorf("grep failed: %w", err)
		}

		if len(matches) == 0 {
			fmt.Fprintln(os.Stderr, "no matches found")
			return nil
		}

		search.PrintText(matches)
		return nil
	},
}

func init() {
	grepCmd.Flags().StringVar(&grepFile, "file", "", "File name glob filter (e.g. '*.json')")
	grepCmd.Flags().BoolVarP(&grepCaseInsensitive, "ignore-case", "i", false, "Case-insensitive search")
	grepCmd.Flags().IntVarP(&grepContext, "context", "C", 0, "Context lines before and after match")
	grepCmd.Flags().StringVar(&grepDir, "dir", "", "Search directory (default: current directory)")
}

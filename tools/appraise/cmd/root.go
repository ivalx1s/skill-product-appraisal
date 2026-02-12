// Package cmd implements the cobra CLI commands for the appraise tool.
package cmd

import (
	"os"

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
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(grepCmd)
	rootCmd.AddCommand(calcCmd)
}

package cmd

import "github.com/spf13/cobra"

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Inspect table metadata",
	Long:  "Commands for locating tables and viewing table schema placeholders.",
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
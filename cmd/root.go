/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var debug bool
var configPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "archery-cli",
	Short: "CLI for Archery API",
	Long:  "A command-line client for Archery API workflows.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Print debug request/response placeholders")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to CLI config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func printDebugRequestResponse(command string, request any, response any) {
	if !debug {
		return
	}

	requestText := formatDebugValue(request)
	responseText := formatDebugValue(response)

	fmt.Printf("[debug] %s request: %s\n", command, requestText)
	fmt.Printf("[debug] %s response: %s\n", command, responseText)
}

func formatDebugValue(v any) string {
	if v == nil {
		return "null"
	}

	body, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}

	return string(body)
}

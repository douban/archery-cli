/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

var submitStatement string
var submitStatementFile string
var submitInstance string
var submitTitle string
var submitRequirementURL string
var submitExtra string

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"request"},
	Short:   "Submit schema change request",
	Long:    "Create and submit table/schema change requests to Archery API.",
	RunE: func(cmd *cobra.Command, args []string) error {
		statement, err := libs.ResolveStatementInput(submitStatement, submitStatementFile)
		if err != nil {
			return err
		}

		extra := map[string]any{}
		if submitExtra != "" {
			if err := json.Unmarshal([]byte(submitExtra), &extra); err != nil {
				return fmt.Errorf("parse --extra as JSON object: %w", err)
			}
		}

		cfg, err := libs.LoadConfig(configPath)
		if err != nil {
			return err
		}

		printDebugRequestResponse("submit", map[string]any{
			"base_url":        cfg.BaseURL,
			"api_token":       "***",
			"statement":       statement,
			"instance":        submitInstance,
			"title":           submitTitle,
			"requirement_url": submitRequirementURL,
			"extra":           extra,
		}, map[string]any{
			"status": "placeholder",
		})

		fmt.Println("submit execution is not implemented yet")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	requestCmd.Flags().StringVar(&submitStatement, "statement", "", "SQL statement text")
	requestCmd.Flags().StringVar(&submitStatementFile, "statement-file", "", "Path to file containing SQL statement")
	requestCmd.Flags().StringVar(&submitInstance, "instance", "", "Database instance")
	requestCmd.Flags().StringVar(&submitTitle, "title", "", "Change request title")
	requestCmd.Flags().StringVar(&submitRequirementURL, "requirement-url", "", "Requirement link URL")
	requestCmd.Flags().StringVar(&submitExtra, "extra", "", "Advanced parameters in JSON object format")

	_ = requestCmd.MarkFlagRequired("instance")
	_ = requestCmd.MarkFlagRequired("title")
}

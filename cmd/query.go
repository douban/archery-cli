/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

var queryStatement string
var queryStatementFile string
var queryDatabase string
var queryInstance string
var querySchema string
var queryTable string
var queryLimitNum int
var queryOutput string

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query statements",
	Long:  "Run read-only query statement workflows against Archery API.",
	RunE: func(cmd *cobra.Command, args []string) error {
		statement, err := libs.ResolveStatementInput(queryStatement, queryStatementFile)
		if err != nil {
			return err
		}

		runtime, err := libs.LoadRuntime(configPath)
		if err != nil {
			return err
		}

		outputFormat, err := libs.ResolveOutputFormat(queryOutput, runtime.Config.DefaultOutput)
		if err != nil {
			return err
		}

		instance := strings.TrimSpace(queryInstance)
		if instance == "" {
			return fmt.Errorf("--instance is required")
		}

		result, err := runtime.ExecuteQuery(
			instance,
			queryDatabase,
			querySchema,
			queryTable,
			statement,
			queryLimitNum,
		)
		if err != nil {
			return err
		}

		envelope := result.Envelope

		printDebugRequestResponse("query", map[string]any{
			"base_url":  runtime.Config.BaseURL,
			"api_token": "***",
			"instance":  instance,
			"database":  queryDatabase,
			"schema":    querySchema,
			"table":     queryTable,
			"statement": statement,
			"limit_num": queryLimitNum,
			"output":    outputFormat,
		}, map[string]any{
			"envelope": envelope,
			"raw":      string(result.RawBody),
		})

		if envelope.Status != 0 {
			if strings.TrimSpace(envelope.Msg) == "" {
				return fmt.Errorf("query failed")
			}
			return fmt.Errorf("%s", envelope.Msg)
		}

		if envelope.Data == nil {
			envelope.Data = map[string]any{}
		}

		if errorMessage, ok := envelope.Data["error"].(string); ok && strings.TrimSpace(errorMessage) != "" {
			return fmt.Errorf("%s", errorMessage)
		}

		headers := anySliceToStrings(envelope.Data["column_list"])
		rows := anyRowsToStrings(envelope.Data["rows"])

		if outputFormat == "text" {
			return libs.RenderText(os.Stdout, formatRowsText(headers, rows))
		}

		if err := libs.RenderTable(os.Stdout, headers, rows); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	queryCmd.Flags().StringVar(&queryStatement, "statement", "", "SQL statement text")
	queryCmd.Flags().StringVar(&queryStatementFile, "statement-file", "", "Path to file containing SQL statement")
	queryCmd.Flags().StringVar(&queryDatabase, "database", "", "Database name")
	queryCmd.Flags().StringVar(&queryInstance, "instance", "", "Database instance")
	queryCmd.Flags().StringVar(&querySchema, "schema", "", "Schema name")
	queryCmd.Flags().StringVar(&queryTable, "table", "", "Table name")
	queryCmd.Flags().IntVar(&queryLimitNum, "limit-num", 0, "Maximum number of rows to return")
	queryCmd.Flags().StringVar(&queryOutput, "output", "", "Output format: table or text")

	_ = queryCmd.MarkFlagRequired("instance")
	_ = queryCmd.MarkFlagRequired("database")
}

func anySliceToStrings(value any) []string {
	if value == nil {
		return []string{}
	}

	switch typed := value.(type) {
	case []string:
		return typed
	case []any:
		result := make([]string, 0, len(typed))
		for _, item := range typed {
			result = append(result, fmt.Sprint(item))
		}
		return result
	default:
		return []string{fmt.Sprint(value)}
	}
}

func anyRowsToStrings(value any) [][]string {
	if value == nil {
		return [][]string{}
	}

	rows, ok := value.([]any)
	if !ok {
		return [][]string{{fmt.Sprint(value)}}
	}

	result := make([][]string, 0, len(rows))
	for _, rowValue := range rows {
		switch row := rowValue.(type) {
		case []any:
			resultRow := make([]string, 0, len(row))
			for _, cell := range row {
				resultRow = append(resultRow, fmt.Sprint(cell))
			}
			result = append(result, resultRow)
		case []string:
			result = append(result, row)
		default:
			result = append(result, []string{fmt.Sprint(rowValue)})
		}
	}

	return result
}

func formatRowsText(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}

	lines := make([]string, 0, len(rows)+1)
	lines = append(lines, strings.Join(headers, "\t"))
	for _, row := range rows {
		lines = append(lines, strings.Join(row, "\t"))
	}

	return strings.Join(lines, "\n")
}

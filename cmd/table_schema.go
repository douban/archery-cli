package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

var tableSchemaInstance string
var tableSchemaDatabase string
var tableSchemaName string

var tableSchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Show table schema",
	Long:  "Display table schema text for a specific instance, database, and table.",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := libs.LoadRuntime(configPath)
		if err != nil {
			return err
		}

		instance := strings.TrimSpace(tableSchemaInstance)
		database := strings.TrimSpace(tableSchemaDatabase)
		tableName := strings.TrimSpace(tableSchemaName)
		result, err := runtime.DescribeTableStructure(instance, database, tableName, "")
		if err != nil {
			return err
		}

		response := result.Envelope

		printDebugRequestResponse("table schema", map[string]any{
			"base_url":  runtime.Config.BaseURL,
			"api_token": "***",
			"instance":  instance,
			"database":  database,
			"table":     tableName,
		}, map[string]any{
			"envelope": response,
			"raw":      string(result.RawBody),
		})

		if response.Status != 0 {
			if strings.TrimSpace(response.Msg) == "" {
				return fmt.Errorf("schema query failed")
			}
			return fmt.Errorf("%s", response.Msg)
		}

		if response.Data == nil {
			return libs.RenderText(os.Stdout, "")
		}

		headers := anySliceToStrings(response.Data["column_list"])
		rows := anyRowsToStrings(response.Data["rows"])
		if len(headers) > 0 && len(rows) > 0 {
			if err := libs.RenderTable(os.Stdout, headers, rows); err != nil {
				return err
			}
			return nil
		}

		if schema, ok := response.Data["create_sql"].(string); ok && strings.TrimSpace(schema) != "" {
			return libs.RenderText(os.Stdout, schema)
		}

		if raw, ok := response.Data["rows"]; ok {
			return libs.RenderText(os.Stdout, fmt.Sprint(raw))
		}

		return libs.RenderText(os.Stdout, fmt.Sprint(response.Data))
	},
}

func init() {
	tableCmd.AddCommand(tableSchemaCmd)
	tableSchemaCmd.Flags().StringVar(&tableSchemaInstance, "instance", "", "Database instance")
	tableSchemaCmd.Flags().StringVar(&tableSchemaDatabase, "database", "", "Database name")
	tableSchemaCmd.Flags().StringVar(&tableSchemaName, "table", "", "Table name")

	_ = tableSchemaCmd.MarkFlagRequired("instance")
	_ = tableSchemaCmd.MarkFlagRequired("database")
	_ = tableSchemaCmd.MarkFlagRequired("table")
}

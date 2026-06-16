package cmd

import (
	"os"
	"strings"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

var tableLocateCmd = &cobra.Command{
	Use:   "locate TABLE_NAME",
	Short: "Locate a table",
	Long:  "Find the instance and database for a table.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := libs.LoadRuntime(configPath)
		if err != nil {
			return err
		}

		tableName := strings.TrimSpace(args[0])
		result, err := runtime.LookupTableInstances(cmd.Context(), tableName)
		if err != nil {
			printDebugRequestResponse("table locate", map[string]any{
				"base_url":  runtime.Config.BaseURL,
				"api_token": runtime.Config.APIToken,
				"table":     tableName,
			}, map[string]any{
				"error": err.Error(),
			})
			return err
		}
		response := result.Response

		rows := make([][]string, 0, len(response.Data))
		for _, item := range response.Data {
			tableNameValue := ""
			if item.HasTableName() {
				tableNameValue = item.GetTableName()
			}
			rows = append(rows, []string{
				item.GetName(),
				item.GetDbName(),
				tableNameValue,
			})
		}
		if debug {
			printDebugRequestResponse("table locate", map[string]any{
				"base_url":  runtime.Config.BaseURL,
				"api_token": "***",
				"table":     tableName,
			}, map[string]any{
				"response": response,
				"raw":      string(result.RawBody),
			})
		}

		if err := libs.RenderTable(os.Stdout, []string{"Instance", "Database", "Table"}, rows); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	tableCmd.AddCommand(tableLocateCmd)
}

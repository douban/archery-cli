/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Archery API",
	Long:  "Clear current authentication state and end the session.",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := libs.RemoveConfig(configPath)
		if err != nil {
			return err
		}

		printDebugRequestResponse("logout", map[string]any{}, map[string]any{
			"removed": path,
		})
		fmt.Printf("removed config at %s\n", path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

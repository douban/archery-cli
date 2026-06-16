/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/douban/archery-cli/libs"
	"github.com/spf13/cobra"
)

var loginBaseURL string
var loginAPIToken string
var loginAPIKeyAlias string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Archery API",
	Long:  "Authenticate and initialize a session for subsequent commands.",
	RunE: func(cmd *cobra.Command, args []string) error {
		token := strings.TrimSpace(loginAPIToken)
		if token == "" {
			token = strings.TrimSpace(loginAPIKeyAlias)
		}
		if token == "" {
			var err error
			token, err = promptLoginToken()
			if err != nil {
				return err
			}
		}

		cfg := libs.Config{
			BaseURL:  strings.TrimSpace(loginBaseURL),
			APIToken: token,
		}

		path, err := libs.SaveConfig(configPath, cfg)
		if err != nil {
			return err
		}

		printDebugRequestResponse("login", map[string]any{
			"base_url":  cfg.BaseURL,
			"api_token": "***",
		}, map[string]any{
			"saved_to": path,
		})

		fmt.Printf("saved config to %s\n", path)
		return nil
	},
}

func promptLoginToken() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter API token: ")
	value, err := reader.ReadString('\n')
	if err != nil && value == "" {
		return "", fmt.Errorf("read api token: %w", err)
	}

	token := strings.TrimSpace(value)
	if token == "" {
		return "", fmt.Errorf("api token cannot be empty")
	}

	return token, nil
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVar(&loginBaseURL, "base-url", "", "Archery API base URL, e.g. https://archery.example.com")
	loginCmd.Flags().StringVar(&loginAPIToken, "api-token", "", "Archery API token (stored in plaintext)")
	loginCmd.Flags().StringVar(&loginAPIKeyAlias, "api-key", "", "Deprecated alias for --api-token")

	_ = loginCmd.MarkFlagRequired("base-url")
}

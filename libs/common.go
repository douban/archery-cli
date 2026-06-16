package libs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const ConfigPathEnv = "ARCHERY_CLI_CONFIG"

type Config struct {
	BaseURL string `json:"base_url"`
	APIToken string `json:"api_token"`
	DefaultOutput string `json:"default_output,omitempty"`
}

func ResolveConfigPath(customPath string) (string, error) {
	if value := strings.TrimSpace(customPath); value != "" {
		return filepath.Abs(value)
	}

	if value := strings.TrimSpace(os.Getenv(ConfigPathEnv)); value != "" {
		return filepath.Abs(value)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve user home: %w", err)
	}

	return filepath.Join(home, ".archery-cli", "config.json"), nil
}

func SaveConfig(customPath string, cfg Config) (string, error) {
	path, err := ResolveConfigPath(customPath)
	if err != nil {
		return "", err
	}

	cfg.BaseURL = strings.TrimSpace(cfg.BaseURL)
	cfg.APIToken = strings.TrimSpace(cfg.APIToken)
	cfg.DefaultOutput = strings.TrimSpace(cfg.DefaultOutput)
	if cfg.DefaultOutput == "" {
		cfg.DefaultOutput = "table"
	}

	if cfg.BaseURL == "" || cfg.APIToken == "" {
		return "", fmt.Errorf("config requires both base_url and api_token")
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return "", fmt.Errorf("create config directory: %w", err)
	}

	body, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal auth config: %w", err)
	}

	// Data is intentionally stored in plaintext per current requirement.
	if err := os.WriteFile(path, body, 0o600); err != nil {
		return "", fmt.Errorf("write auth config: %w", err)
	}

	return path, nil
}

func LoadConfig(customPath string) (Config, error) {
	path, err := ResolveConfigPath(customPath)
	if err != nil {
		return Config{}, err
	}

	body, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, fmt.Errorf("config not found, run 'archery-cli login' first")
		}
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if strings.TrimSpace(cfg.BaseURL) == "" || strings.TrimSpace(cfg.APIToken) == "" {
		return Config{}, fmt.Errorf("config is incomplete, run 'archery-cli login' again")
	}

	if strings.TrimSpace(cfg.DefaultOutput) == "" {
		cfg.DefaultOutput = "table"
	}

	return cfg, nil
}

func RemoveConfig(customPath string) (string, error) {
	path, err := ResolveConfigPath(customPath)
	if err != nil {
		return "", err
	}

	if err := os.Remove(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return path, nil
		}
		return "", fmt.Errorf("remove config: %w", err)
	}

	return path, nil
}

func ResolveStatementInput(statement string, statementFile string) (string, error) {
	hasStatement := strings.TrimSpace(statement) != ""
	hasStatementFile := strings.TrimSpace(statementFile) != ""

	if hasStatement && hasStatementFile {
		return "", fmt.Errorf("use only one of --statement or --statement-file")
	}

	if !hasStatement && !hasStatementFile {
		return "", fmt.Errorf("either --statement or --statement-file is required")
	}

	if hasStatement {
		return statement, nil
	}

	body, err := os.ReadFile(statementFile)
	if err != nil {
		return "", fmt.Errorf("read statement file: %w", err)
	}

	value := strings.TrimSpace(string(body))
	if value == "" {
		return "", fmt.Errorf("statement file is empty")
	}

	return value, nil
}
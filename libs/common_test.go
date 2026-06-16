package libs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveConfigPathPrefersExplicitPath(t *testing.T) {
	t.Setenv(ConfigPathEnv, filepath.Join(t.TempDir(), "env.json"))

	path, err := ResolveConfigPath("./testdata/config.json")
	if err != nil {
		t.Fatalf("ResolveConfigPath returned error: %v", err)
	}

	if !filepath.IsAbs(path) {
		t.Fatalf("expected absolute path, got %q", path)
	}
	if filepath.Base(path) != "config.json" {
		t.Fatalf("expected config.json path, got %q", path)
	}
}

func TestSaveLoadAndRemoveConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	writtenPath, err := SaveConfig(path, Config{
		BaseURL:  "https://archery.example.com",
		APIToken: "token-123",
	})
	if err != nil {
		t.Fatalf("SaveConfig returned error: %v", err)
	}

	loaded, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig returned error: %v", err)
	}

	if loaded.BaseURL != "https://archery.example.com" {
		t.Fatalf("unexpected base url: %q", loaded.BaseURL)
	}
	if loaded.APIToken != "token-123" {
		t.Fatalf("unexpected api token: %q", loaded.APIToken)
	}
	if loaded.DefaultOutput != "table" {
		t.Fatalf("expected default output table, got %q", loaded.DefaultOutput)
	}

	removedPath, err := RemoveConfig(path)
	if err != nil {
		t.Fatalf("RemoveConfig returned error: %v", err)
	}
	if writtenPath != removedPath {
		t.Fatalf("expected removed path %q, got %q", writtenPath, removedPath)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected config file to be removed, stat err=%v", err)
	}
}

func TestResolveStatementInput(t *testing.T) {
	t.Run("inline statement", func(t *testing.T) {
		statement, err := ResolveStatementInput("select 1", "")
		if err != nil {
			t.Fatalf("ResolveStatementInput returned error: %v", err)
		}
		if statement != "select 1" {
			t.Fatalf("unexpected statement: %q", statement)
		}
	})

	t.Run("statement file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "query.sql")
		if err := os.WriteFile(path, []byte("\nselect * from users;\n"), 0o600); err != nil {
			t.Fatalf("WriteFile returned error: %v", err)
		}

		statement, err := ResolveStatementInput("", path)
		if err != nil {
			t.Fatalf("ResolveStatementInput returned error: %v", err)
		}
		if statement != "select * from users;" {
			t.Fatalf("unexpected statement: %q", statement)
		}
	})

	t.Run("mutually exclusive flags", func(t *testing.T) {
		_, err := ResolveStatementInput("select 1", "query.sql")
		if err == nil {
			t.Fatal("expected error for conflicting statement inputs")
		}
	})
}
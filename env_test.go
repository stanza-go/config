package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotenv(t *testing.T) {
	t.Run("parses key-value pairs", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		err := os.WriteFile(envPath, []byte("TEST_KEY=test_value\n"), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		// Clean up env var after test
		defer os.Unsetenv("TEST_KEY")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("TEST_KEY"); got != "test_value" {
			t.Errorf("TEST_KEY = %q, want %q", got, "test_value")
		}
	})

	t.Run("handles double-quoted values", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		err := os.WriteFile(envPath, []byte(`QUOTED_KEY="hello world"`+"\n"), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		defer os.Unsetenv("QUOTED_KEY")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("QUOTED_KEY"); got != "hello world" {
			t.Errorf("QUOTED_KEY = %q, want %q", got, "hello world")
		}
	})

	t.Run("handles single-quoted values", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		err := os.WriteFile(envPath, []byte(`SINGLE_KEY='single quoted'`+"\n"), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		defer os.Unsetenv("SINGLE_KEY")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("SINGLE_KEY"); got != "single quoted" {
			t.Errorf("SINGLE_KEY = %q, want %q", got, "single quoted")
		}
	})

	t.Run("skips comments and empty lines", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		content := `# This is a comment
VALID_KEY=valid_value

# Another comment
ANOTHER_KEY=another_value
`
		err := os.WriteFile(envPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		defer os.Unsetenv("VALID_KEY")
		defer os.Unsetenv("ANOTHER_KEY")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("VALID_KEY"); got != "valid_value" {
			t.Errorf("VALID_KEY = %q, want %q", got, "valid_value")
		}
		if got := os.Getenv("ANOTHER_KEY"); got != "another_value" {
			t.Errorf("ANOTHER_KEY = %q, want %q", got, "another_value")
		}
	})

	t.Run("skips lines without equals sign", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		content := `INVALID_LINE_NO_EQUALS
VALID_LINE=has_value
`
		err := os.WriteFile(envPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		defer os.Unsetenv("VALID_LINE")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("VALID_LINE"); got != "has_value" {
			t.Errorf("VALID_LINE = %q, want %q", got, "has_value")
		}
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		err := loadDotenv("/non/existent/path/.env")
		if err == nil {
			t.Error("expected error for non-existent file, got nil")
		}
	})

	t.Run("handles values with spaces around equals", func(t *testing.T) {
		tempDir := t.TempDir()
		envPath := filepath.Join(tempDir, ".env")
		err := os.WriteFile(envPath, []byte("SPACED_KEY = spaced_value\n"), 0644)
		if err != nil {
			t.Fatalf("failed to create .env file: %v", err)
		}

		defer os.Unsetenv("SPACED_KEY")

		err = loadDotenv(envPath)
		if err != nil {
			t.Fatalf("loadDotenv returned error: %v", err)
		}

		if got := os.Getenv("SPACED_KEY"); got != "spaced_value" {
			t.Errorf("SPACED_KEY = %q, want %q", got, "spaced_value")
		}
	})
}

func TestGetEnvValue(t *testing.T) {
	t.Run("converts dot notation to uppercase underscore", func(t *testing.T) {
		os.Setenv("DB_HOST", "localhost")
		defer os.Unsetenv("DB_HOST")

		val, ok := getEnvValue("db.host")
		if !ok {
			t.Error("expected ok=true, got false")
		}
		if val != "localhost" {
			t.Errorf("getEnvValue(db.host) = %q, want %q", val, "localhost")
		}
	})

	t.Run("returns false for non-existent env var", func(t *testing.T) {
		_, ok := getEnvValue("non.existent.key")
		if ok {
			t.Error("expected ok=false for non-existent key, got true")
		}
	})

	t.Run("handles nested keys", func(t *testing.T) {
		os.Setenv("SERVICES_THRIFT_TIMEOUT", "30")
		defer os.Unsetenv("SERVICES_THRIFT_TIMEOUT")

		val, ok := getEnvValue("services.thrift.timeout")
		if !ok {
			t.Error("expected ok=true, got false")
		}
		if val != "30" {
			t.Errorf("getEnvValue(services.thrift.timeout) = %q, want %q", val, "30")
		}
	})
}

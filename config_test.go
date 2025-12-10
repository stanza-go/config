package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Create a temporary config file for testing
		tempDir := t.TempDir()
		configPath := filepath.Join(tempDir, "config.yaml")
		err := os.WriteFile(configPath, []byte("app:\n  env: test\n"), 0644)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Change to temp directory
		originalDir, _ := os.Getwd()
		err = os.Chdir(tempDir)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer os.Chdir(originalDir)

		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("unexpected panic: %v", r)
				}
			}()
			Init()
		}()
	})
}

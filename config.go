package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// configData holds the parsed YAML configuration
var (
	configData map[string]any
	mu         sync.RWMutex
)

// Init initializes the configuration by loading .env file (if exists) and config.yaml
func Init() {
	path := lookupConfigPath()

	// Load .env file if it exists
	envPath := filepath.Join(path, ".env")
	if _, err := os.Stat(envPath); err == nil {
		if err := loadDotenv(envPath); err != nil {
			log.Printf("loadDotenv returns error: %s\n", err.Error())
		}
	}

	// Load config.yaml
	configPath := filepath.Join(path, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("os.ReadFile returns error: %s\n", err.Error())
	}

	mu.Lock()
	defer mu.Unlock()

	if err := yaml.Unmarshal(data, &configData); err != nil {
		log.Fatalf("yaml.Unmarshal returns error: %s\n", err.Error())
	}
}

// lookupConfigPath searches for config.yaml starting from the current directory
// and traversing up to parent directories
func lookupConfigPath() string {
	pathHasConfigFile := func(path string) bool {
		filePath := filepath.Join(path, "config.yaml")

		_, err := os.Stat(filePath)
		if err == nil {
			return true
		} else if os.IsNotExist(err) {
			return false
		} else {
			log.Fatalf("os.Stat returns error: %s (%s)", err.Error(), filePath)
			return false
		}
	}

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd returns error: %s (%s)", err.Error(), currentPath)
	}

	for {
		if pathHasConfigFile(currentPath) {
			return currentPath
		}

		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			return currentPath
		}

		currentPath = parentPath
	}
}

// getFromMap retrieves a value from a nested map using dot notation (e.g., "db.host")
func getFromMap(key string) (any, bool) {
	mu.RLock()
	defer mu.RUnlock()

	parts := strings.Split(key, ".")
	var current any = configData

	for _, part := range parts {
		m, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		current, ok = m[part]
		if !ok {
			return nil, false
		}
	}
	return current, true
}

// setInMap sets a value in the nested map using dot notation (e.g., "db.host")
// This is primarily used for testing
func setInMap(key string, value any) {
	mu.Lock()
	defer mu.Unlock()

	if configData == nil {
		configData = make(map[string]any)
	}

	parts := strings.Split(key, ".")
	current := configData

	for i, part := range parts {
		if i == len(parts)-1 {
			current[part] = value
			return
		}

		if _, ok := current[part]; !ok {
			current[part] = make(map[string]any)
		}

		if next, ok := current[part].(map[string]any); ok {
			current = next
		} else {
			current[part] = make(map[string]any)
			current = current[part].(map[string]any)
		}
	}
}

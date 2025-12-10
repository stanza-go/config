package config

import (
	"bufio"
	"os"
	"strings"
)

// loadDotenv parses a .env file and sets environment variables
func loadDotenv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		idx := strings.Index(line, "=")
		if idx == -1 {
			continue
		}

		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+1:])

		// Remove surrounding quotes if present
		if len(value) >= 2 {
			if (value[0] == '"' && value[len(value)-1] == '"') ||
				(value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			}
		}

		os.Setenv(key, value)
	}

	return scanner.Err()
}

// getEnvValue checks for an environment variable override
// Converts "db.host" -> "DB_HOST"
func getEnvValue(key string) (string, bool) {
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	return os.LookupEnv(envKey)
}

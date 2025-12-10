package config

import (
	"time"

	"gopkg.in/yaml.v3"
)

// IsSet reports whether the given key exists in the configuration.
//
// A key is considered "set" if it exists in either:
//  1. Environment variables (key converted to UPPER_SNAKE_CASE)
//  2. The config file (config.yaml)
//
// Config file example (config.yaml):
//
//	app:
//	  name: myapp
//
// Usage:
//
//	if config.IsSet("app.name") {
//	    name := config.GetString("app.name")
//	}
func IsSet(key string) bool {
	if _, ok := getEnvValue(key); ok {
		return true
	}
	_, ok := getFromMap(key)
	return ok
}

// GetString returns the string value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE, e.g., "app.name" -> "APP_NAME")
//  2. Config file value
//
// Returns an empty string if the key is not found.
//
// Config file example (config.yaml):
//
//	app:
//	  name: myapp
//	  env: production
//
// Environment variable override:
//
//	export APP_NAME=otherapp
//
// Usage:
//
//	name := config.GetString("app.name")  // returns "myapp", or "otherapp" if env var is set
func GetString(key string) string {
	if val, ok := getEnvValue(key); ok {
		return val
	}
	if val, ok := getFromMap(key); ok {
		return toString(val)
	}
	return ""
}

// GetBool returns the boolean value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE, e.g., "app.debug" -> "APP_DEBUG")
//  2. Config file value
//
// Returns false if the key is not found or cannot be converted to bool.
//
// Accepted truthy values: true, "true", "1"
// Accepted falsy values: false, "false", "0"
//
// Config file example (config.yaml):
//
//	app:
//	  debug: true
//
// Environment variable override:
//
//	export APP_DEBUG=false
//
// Usage:
//
//	if config.GetBool("app.debug") {
//	    log.SetLevel(log.DebugLevel)
//	}
func GetBool(key string) bool {
	if val, ok := getEnvValue(key); ok {
		return toBool(val)
	}
	if val, ok := getFromMap(key); ok {
		return toBool(val)
	}
	return false
}

// GetInt returns the integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE, e.g., "http.port" -> "HTTP_PORT")
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to int.
//
// Config file example (config.yaml):
//
//	http:
//	  port: 8080
//
// Environment variable override:
//
//	export HTTP_PORT=3000
//
// Usage:
//
//	port := config.GetInt("http.port")  // returns 8080, or 3000 if env var is set
func GetInt(key string) int {
	if val, ok := getEnvValue(key); ok {
		return toInt(val)
	}
	if val, ok := getFromMap(key); ok {
		return toInt(val)
	}
	return 0
}

// GetInt32 returns the 32-bit integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to int32.
//
// Config file example (config.yaml):
//
//	limits:
//	  max_connections: 1000
//
// Usage:
//
//	maxConn := config.GetInt32("limits.max_connections")
func GetInt32(key string) int32 {
	if val, ok := getEnvValue(key); ok {
		return toInt32(val)
	}
	if val, ok := getFromMap(key); ok {
		return toInt32(val)
	}
	return 0
}

// GetInt64 returns the 64-bit integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to int64.
//
// Config file example (config.yaml):
//
//	storage:
//	  max_file_size: 10737418240  # 10GB in bytes
//
// Usage:
//
//	maxSize := config.GetInt64("storage.max_file_size")
func GetInt64(key string) int64 {
	if val, ok := getEnvValue(key); ok {
		return toInt64(val)
	}
	if val, ok := getFromMap(key); ok {
		return toInt64(val)
	}
	return 0
}

// GetUint returns the unsigned integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to uint.
//
// Config file example (config.yaml):
//
//	worker:
//	  pool_size: 10
//
// Usage:
//
//	poolSize := config.GetUint("worker.pool_size")
func GetUint(key string) uint {
	if val, ok := getEnvValue(key); ok {
		return toUint(val)
	}
	if val, ok := getFromMap(key); ok {
		return toUint(val)
	}
	return 0
}

// GetUint16 returns the 16-bit unsigned integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to uint16.
//
// Config file example (config.yaml):
//
//	http:
//	  port: 8080
//
// Usage:
//
//	port := config.GetUint16("http.port")
func GetUint16(key string) uint16 {
	if val, ok := getEnvValue(key); ok {
		return toUint16(val)
	}
	if val, ok := getFromMap(key); ok {
		return toUint16(val)
	}
	return 0
}

// GetUint32 returns the 32-bit unsigned integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to uint32.
//
// Config file example (config.yaml):
//
//	cache:
//	  max_items: 100000
//
// Usage:
//
//	maxItems := config.GetUint32("cache.max_items")
func GetUint32(key string) uint32 {
	if val, ok := getEnvValue(key); ok {
		return toUint32(val)
	}
	if val, ok := getFromMap(key); ok {
		return toUint32(val)
	}
	return 0
}

// GetUint64 returns the 64-bit unsigned integer value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to uint64.
//
// Config file example (config.yaml):
//
//	metrics:
//	  max_data_points: 18446744073709551615
//
// Usage:
//
//	maxPoints := config.GetUint64("metrics.max_data_points")
func GetUint64(key string) uint64 {
	if val, ok := getEnvValue(key); ok {
		return toUint64(val)
	}
	if val, ok := getFromMap(key); ok {
		return toUint64(val)
	}
	return 0
}

// GetFloat64 returns the float64 value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to float64.
//
// Config file example (config.yaml):
//
//	ml:
//	  learning_rate: 0.001
//	  threshold: 0.85
//
// Environment variable override:
//
//	export ML_LEARNING_RATE=0.01
//
// Usage:
//
//	rate := config.GetFloat64("ml.learning_rate")
func GetFloat64(key string) float64 {
	if val, ok := getEnvValue(key); ok {
		return toFloat64(val)
	}
	if val, ok := getFromMap(key); ok {
		return toFloat64(val)
	}
	return 0
}

// GetDuration returns the [time.Duration] value associated with the given key.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns 0 if the key is not found or cannot be converted to duration.
//
// Supported formats:
//   - Duration strings: "300ms", "1.5s", "2m", "1h30m", "24h"
//   - Integer values: treated as nanoseconds
//
// Config file example (config.yaml):
//
//	http:
//	  read_timeout: 30s
//	  write_timeout: 1m
//	  idle_timeout: 2h
//
// Environment variable override:
//
//	export HTTP_READ_TIMEOUT=60s
//
// Usage:
//
//	timeout := config.GetDuration("http.read_timeout")
//	server := &http.Server{
//	    ReadTimeout: timeout,
//	}
func GetDuration(key string) time.Duration {
	if val, ok := getEnvValue(key); ok {
		return toDuration(val)
	}
	if val, ok := getFromMap(key); ok {
		return toDuration(val)
	}
	return 0
}

// GetStringPtr returns a pointer to the string value associated with the given key.
//
// Returns nil if the key is not set (does not exist in config or environment).
// This is useful for distinguishing between an empty string value and a missing key.
//
// Config file example (config.yaml):
//
//	database:
//	  password: secret123
//
// Usage:
//
//	password := config.GetStringPtr("database.password")
//	if password == nil {
//	    log.Fatal("database password is required")
//	}
//	db.Connect(*password)
func GetStringPtr(key string) *string {
	if !IsSet(key) {
		return nil
	}
	value := GetString(key)
	return &value
}

// GetStringOr returns the string value associated with the given key,
// or the defaultValue if the key is not set.
//
// Config file example (config.yaml):
//
//	app:
//	  env: production
//
// Usage:
//
//	env := config.GetStringOr("app.env", "development")  // returns "production"
//	region := config.GetStringOr("app.region", "us-east-1")  // returns "us-east-1" (not in config)
func GetStringOr(key string, defaultValue string) string {
	if IsSet(key) {
		return GetString(key)
	}
	return defaultValue
}

// GetBoolOr returns the boolean value associated with the given key,
// or the defaultValue if the key is not set.
//
// Config file example (config.yaml):
//
//	features:
//	  dark_mode: true
//
// Usage:
//
//	darkMode := config.GetBoolOr("features.dark_mode", false)
//	analytics := config.GetBoolOr("features.analytics", true)  // defaults to true if not set
func GetBoolOr(key string, defaultValue bool) bool {
	if IsSet(key) {
		return GetBool(key)
	}
	return defaultValue
}

// GetIntOr returns the integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Config file example (config.yaml):
//
//	http:
//	  port: 8080
//
// Usage:
//
//	port := config.GetIntOr("http.port", 3000)  // returns 8080
//	workers := config.GetIntOr("http.workers", 4)  // returns 4 (not in config)
func GetIntOr(key string, defaultValue int) int {
	if IsSet(key) {
		return GetInt(key)
	}
	return defaultValue
}

// GetInt32Or returns the 32-bit integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	maxConn := config.GetInt32Or("db.max_connections", 100)
func GetInt32Or(key string, defaultValue int32) int32 {
	if IsSet(key) {
		return GetInt32(key)
	}
	return defaultValue
}

// GetInt64Or returns the 64-bit integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	maxSize := config.GetInt64Or("upload.max_size", 10*1024*1024)  // default 10MB
func GetInt64Or(key string, defaultValue int64) int64 {
	if IsSet(key) {
		return GetInt64(key)
	}
	return defaultValue
}

// GetUintOr returns the unsigned integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	poolSize := config.GetUintOr("worker.pool_size", 5)
func GetUintOr(key string, defaultValue uint) uint {
	if IsSet(key) {
		return GetUint(key)
	}
	return defaultValue
}

// GetUint16Or returns the 16-bit unsigned integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	port := config.GetUint16Or("grpc.port", 50051)
func GetUint16Or(key string, defaultValue uint16) uint16 {
	if IsSet(key) {
		return GetUint16(key)
	}
	return defaultValue
}

// GetUint32Or returns the 32-bit unsigned integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	bufferSize := config.GetUint32Or("io.buffer_size", 4096)
func GetUint32Or(key string, defaultValue uint32) uint32 {
	if IsSet(key) {
		return GetUint32(key)
	}
	return defaultValue
}

// GetUint64Or returns the 64-bit unsigned integer value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	maxMemory := config.GetUint64Or("cache.max_memory", 1<<30)  // default 1GB
func GetUint64Or(key string, defaultValue uint64) uint64 {
	if IsSet(key) {
		return GetUint64(key)
	}
	return defaultValue
}

// GetFloat64Or returns the float64 value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	rate := config.GetFloat64Or("rate_limiter.requests_per_second", 100.0)
func GetFloat64Or(key string, defaultValue float64) float64 {
	if IsSet(key) {
		return GetFloat64(key)
	}
	return defaultValue
}

// GetDurationOr returns the [time.Duration] value associated with the given key,
// or the defaultValue if the key is not set.
//
// Usage:
//
//	timeout := config.GetDurationOr("http.timeout", 30*time.Second)
//	cacheTTL := config.GetDurationOr("cache.ttl", 5*time.Minute)
func GetDurationOr(key string, defaultValue time.Duration) time.Duration {
	if IsSet(key) {
		return GetDuration(key)
	}
	return defaultValue
}

// GetStringSlice returns a string slice value associated with the given key.
//
// Returns nil if the key is not found.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE, comma-separated)
//  2. Config file value
//
// Config file example (config.yaml):
//
//	cors:
//	  allowed_origins:
//	    - https://example.com
//	    - https://api.example.com
//
// Environment variable override (comma-separated):
//
//	export CORS_ALLOWED_ORIGINS=https://example.com,https://api.example.com
//
// Whitespace around commas is trimmed:
//
//	export TAGS=foo, bar, baz  # returns ["foo", "bar", "baz"]
//
// Empty values are filtered out:
//
//	export TAGS=foo,,bar  # returns ["foo", "bar"]
//
// Usage:
//
//	origins := config.GetStringSlice("cors.allowed_origins")
//	for _, origin := range origins {
//	    fmt.Println(origin)
//	}
func GetStringSlice(key string) []string {
	if val, ok := getEnvValue(key); ok {
		return splitAndTrimStringSlice(val)
	}
	if val, ok := getFromMap(key); ok {
		return toStringSlice(val)
	}
	return nil
}

// GetIntSlice returns an integer slice value associated with the given key.
//
// Returns nil if the key is not found.
//
// Lookup order:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE, comma-separated)
//  2. Config file value
//
// Config file example (config.yaml):
//
//	retry:
//	  backoff_ms:
//	    - 100
//	    - 200
//	    - 500
//	    - 1000
//
// Environment variable override (comma-separated):
//
//	export RETRY_BACKOFF_MS=100,200,500,1000
//
// Whitespace around commas is trimmed:
//
//	export PORTS=8080, 8081, 8082  # returns [8080, 8081, 8082]
//
// Empty values and invalid integers are filtered out:
//
//	export PORTS=8080,,invalid,8081  # returns [8080, 8081]
//
// Usage:
//
//	backoffs := config.GetIntSlice("retry.backoff_ms")
//	for _, ms := range backoffs {
//	    time.Sleep(time.Duration(ms) * time.Millisecond)
//	}
func GetIntSlice(key string) []int {
	if val, ok := getEnvValue(key); ok {
		return splitAndTrimIntSlice(val)
	}
	if val, ok := getFromMap(key); ok {
		return toIntSlice(val)
	}
	return nil
}

// GetStringMap returns a map[string]any value associated with the given key.
//
// Returns an empty map if the key is not found.
//
// Note: This function does not support environment variable override.
//
// Config file example (config.yaml):
//
//	database:
//	  postgres:
//	    host: localhost
//	    port: 5432
//	    name: mydb
//
// Usage:
//
//	dbConfig := config.GetStringMap("database.postgres")
//	host := dbConfig["host"].(string)
//	port := dbConfig["port"].(int)
func GetStringMap(key string) map[string]any {
	if val, ok := getFromMap(key); ok {
		return toStringMap(val)
	}
	return map[string]any{}
}

// Unmarshal unmarshals the entire configuration into the provided struct.
//
// The struct should use `yaml` struct tags to map configuration keys to fields.
// This is useful for loading the entire configuration at once into a typed struct.
//
// Lookup order for each value:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Config file example (config.yaml):
//
//	app:
//	  name: myapp
//	  env: production
//	http:
//	  port: 8080
//
// Usage:
//
//	type Config struct {
//	    App struct {
//	        Name string `yaml:"name"`
//	        Env  string `yaml:"env"`
//	    } `yaml:"app"`
//	    HTTP struct {
//	        Port int `yaml:"port"`
//	    } `yaml:"http"`
//	}
//
//	var cfg Config
//	if err := config.Unmarshal(&cfg); err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("App: %s, Port: %d\n", cfg.App.Name, cfg.HTTP.Port)
func Unmarshal(v any) error {
	mu.RLock()
	defer mu.RUnlock()

	// Apply environment variable overrides before unmarshaling
	configWithOverrides := applyEnvOverrides(configData, "")

	data, err := yaml.Marshal(configWithOverrides)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// UnmarshalKey unmarshals a specific configuration section into the provided struct.
//
// The struct should use `yaml` struct tags to map configuration keys to fields.
// This is useful for loading only a portion of the configuration.
//
// Lookup order for each value:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Returns nil error if the key does not exist (struct remains unchanged).
//
// Config file example (config.yaml):
//
//	database:
//	  host: localhost
//	  port: 5432
//	  name: mydb
//	  max_connections: 100
//
// Usage:
//
//	type DatabaseConfig struct {
//	    Host           string `yaml:"host"`
//	    Port           int    `yaml:"port"`
//	    Name           string `yaml:"name"`
//	    MaxConnections int    `yaml:"max_connections"`
//	}
//
//	var dbCfg DatabaseConfig
//	if err := config.UnmarshalKey("database", &dbCfg); err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Connecting to %s:%d/%s\n", dbCfg.Host, dbCfg.Port, dbCfg.Name)
func UnmarshalKey(key string, v any) error {
	val, ok := getFromMap(key)
	if !ok {
		return nil
	}

	// Apply environment variable overrides if val is a map
	var dataToMarshal any
	if m, ok := val.(map[string]any); ok {
		dataToMarshal = applyEnvOverrides(m, key)
	} else {
		// For non-map values, check for env override
		if envVal, ok := getEnvValue(key); ok {
			dataToMarshal = envVal
		} else {
			dataToMarshal = val
		}
	}

	data, err := yaml.Marshal(dataToMarshal)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// AllSettings returns a copy of all configuration settings as a map.
//
// The returned map is a deep copy with environment variable overrides applied.
// Modifications to the returned map will not affect the underlying configuration.
//
// Lookup order for each value:
//  1. Environment variable (key converted to UPPER_SNAKE_CASE)
//  2. Config file value
//
// Usage:
//
//	settings := config.AllSettings()
//	for key, value := range settings {
//	    fmt.Printf("%s: %v\n", key, value)
//	}
func AllSettings() map[string]any {
	mu.RLock()
	defer mu.RUnlock()

	return applyEnvOverrides(configData, "")
}

// applyEnvOverrides recursively applies environment variable overrides to a map
func applyEnvOverrides(data map[string]any, prefix string) map[string]any {
	result := make(map[string]any)

	for k, v := range data {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}

		switch val := v.(type) {
		case map[string]any:
			// Recursively process nested maps
			result[k] = applyEnvOverrides(val, key)
		default:
			// Check for environment variable override
			if envVal, ok := getEnvValue(key); ok {
				// Convert env string to match original value's type
				result[k] = convertEnvToType(envVal, v)
			} else {
				result[k] = v
			}
		}
	}

	return result
}

// convertEnvToType converts an environment variable string to match the original value's type
func convertEnvToType(envVal string, originalVal any) any {
	switch originalVal.(type) {
	case bool:
		return toBool(envVal)
	case int:
		return toInt(envVal)
	case int32:
		return toInt32(envVal)
	case int64:
		return toInt64(envVal)
	case uint:
		return toUint(envVal)
	case uint16:
		return toUint16(envVal)
	case uint32:
		return toUint32(envVal)
	case uint64:
		return toUint64(envVal)
	case float32, float64:
		return toFloat64(envVal)
	case []string:
		return splitAndTrimStringSlice(envVal)
	case []int:
		return splitAndTrimIntSlice(envVal)
	case []any:
		// For []any from YAML, try to infer type from first element
		slice := originalVal.([]any)
		if len(slice) > 0 {
			switch slice[0].(type) {
			case int, int64, float64:
				return splitAndTrimIntSlice(envVal)
			default:
				return splitAndTrimStringSlice(envVal)
			}
		}
		return splitAndTrimStringSlice(envVal)
	default:
		// For strings and unknown types, return as-is
		return envVal
	}
}

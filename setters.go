package config

// Set stores a value in the in-memory configuration map using dot notation.
//
// This function is primarily intended for testing purposes or programmatic
// configuration. Values set via this function will be overridden by
// environment variables when retrieved using getter functions.
//
// The key uses dot notation to represent nested structure:
//
//	config.Set("database.host", "localhost")
//	config.Set("database.port", 5432)
//
// This is equivalent to the following config.yaml:
//
//	database:
//	  host: localhost
//	  port: 5432
//
// Usage:
//
//	config.Set("app.name", "myapp")
//	config.Set("app.debug", true)
//	config.Set("http.port", 8080)
//
//	name := config.GetString("app.name")  // returns "myapp"
func Set(key string, value any) {
	setInMap(key, value)
}

// Reset clears all configuration data from memory.
//
// This function removes all key-value pairs that were loaded from config.yaml
// or set programmatically via [Set]. It does not affect environment variables.
//
// This is primarily intended for testing purposes to ensure a clean state
// between test cases.
//
// Usage:
//
//	func TestSomething(t *testing.T) {
//	    config.Reset()  // start with clean state
//	    config.Set("app.env", "test")
//	    // ... run test assertions
//	}
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	configData = make(map[string]any)
}

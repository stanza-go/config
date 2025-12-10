package config

import (
	"testing"
)

func TestSetters(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		Reset()

		Set("app.env", "local")
		if got := GetString("app.env"); got != "local" {
			t.Errorf("GetString(app.env) = %v, want %v", got, "local")
		}

		Set("app.debug", true)
		if got := GetBool("app.debug"); got != true {
			t.Errorf("GetBool(app.debug) = %v, want %v", got, true)
		}

		Set("http.port", 8080)
		if got := GetInt("http.port"); got != 8080 {
			t.Errorf("GetInt(http.port) = %v, want %v", got, 8080)
		}
	})

	t.Run("Set nested keys", func(t *testing.T) {
		Reset()

		Set("database.connection.host", "localhost")
		Set("database.connection.port", 5432)

		if got := GetString("database.connection.host"); got != "localhost" {
			t.Errorf("GetString(database.connection.host) = %v, want %v", got, "localhost")
		}
		if got := GetInt("database.connection.port"); got != 5432 {
			t.Errorf("GetInt(database.connection.port) = %v, want %v", got, 5432)
		}
	})

	t.Run("Reset", func(t *testing.T) {
		Reset()
		Set("app.env", "local")
		if got := GetString("app.env"); got != "local" {
			t.Errorf("GetString(app.env) = %v, want %v", got, "local")
		}

		Reset()
		if got := GetString("app.env"); got != "" {
			t.Errorf("GetString(app.env) after Reset = %v, want %v", got, "")
		}
		if IsSet("app.env") {
			t.Error("IsSet(app.env) after Reset = true, want false")
		}
	})
}

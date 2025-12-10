package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGetters(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		Reset()
		Set("app.env", "local")

		if got := GetString("app.env"); got != "local" {
			t.Errorf("GetString(app.env) = %v, want %v", got, "local")
		}
		if got := GetString("app.environment"); got != "" {
			t.Errorf("GetString(app.environment) = %v, want %v", got, "")
		}
	})

	t.Run("GetStringOr", func(t *testing.T) {
		Reset()
		Set("app.env", "local")

		if got := GetStringOr("app.env", "dev"); got != "local" {
			t.Errorf("GetStringOr(app.env, dev) = %v, want %v", got, "local")
		}
		if got := GetStringOr("app.environment", "dev"); got != "dev" {
			t.Errorf("GetStringOr(app.environment, dev) = %v, want %v", got, "dev")
		}
	})

	t.Run("GetBool", func(t *testing.T) {
		Reset()
		Set("app.debug", true)

		if got := GetBool("app.debug"); got != true {
			t.Errorf("GetBool(app.debug) = %v, want %v", got, true)
		}
		if got := GetBool("app.is_debug"); got != false {
			t.Errorf("GetBool(app.is_debug) = %v, want %v", got, false)
		}
	})

	t.Run("GetBoolOr", func(t *testing.T) {
		Reset()
		Set("app.debug", true)

		if got := GetBoolOr("app.debug", false); got != true {
			t.Errorf("GetBoolOr(app.debug, false) = %v, want %v", got, true)
		}
		if got := GetBoolOr("app.is_debug", true); got != true {
			t.Errorf("GetBoolOr(app.is_debug, true) = %v, want %v", got, true)
		}
	})

	t.Run("GetInt", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetInt("http.port"); got != 8080 {
			t.Errorf("GetInt(http.port) = %v, want %v", got, 8080)
		}
		if got := GetInt("https.port"); got != 0 {
			t.Errorf("GetInt(https.port) = %v, want %v", got, 0)
		}
	})

	t.Run("GetIntOr", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetIntOr("http.port", 3000); got != 8080 {
			t.Errorf("GetIntOr(http.port, 3000) = %v, want %v", got, 8080)
		}
		if got := GetIntOr("https.port", 8443); got != 8443 {
			t.Errorf("GetIntOr(https.port, 8443) = %v, want %v", got, 8443)
		}
	})

	t.Run("GetInt32", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetInt32("http.port"); got != int32(8080) {
			t.Errorf("GetInt32(http.port) = %v, want %v", got, int32(8080))
		}
		if got := GetInt32("https.port"); got != int32(0) {
			t.Errorf("GetInt32(https.port) = %v, want %v", got, int32(0))
		}
	})

	t.Run("GetInt32Or", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetInt32Or("http.port", 3000); got != int32(8080) {
			t.Errorf("GetInt32Or(http.port, 3000) = %v, want %v", got, int32(8080))
		}
		if got := GetInt32Or("https.port", 8443); got != int32(8443) {
			t.Errorf("GetInt32Or(https.port, 8443) = %v, want %v", got, int32(8443))
		}
	})

	t.Run("GetInt64", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetInt64("http.port"); got != int64(8080) {
			t.Errorf("GetInt64(http.port) = %v, want %v", got, int64(8080))
		}
		if got := GetInt64("https.port"); got != int64(0) {
			t.Errorf("GetInt64(https.port) = %v, want %v", got, int64(0))
		}
	})

	t.Run("GetInt64Or", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetInt64Or("http.port", 3000); got != int64(8080) {
			t.Errorf("GetInt64Or(http.port, 3000) = %v, want %v", got, int64(8080))
		}
		if got := GetInt64Or("https.port", 8443); got != int64(8443) {
			t.Errorf("GetInt64Or(https.port, 8443) = %v, want %v", got, int64(8443))
		}
	})

	t.Run("GetUint", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint("http.port"); got != uint(8080) {
			t.Errorf("GetUint(http.port) = %v, want %v", got, uint(8080))
		}
		if got := GetUint("https.port"); got != uint(0) {
			t.Errorf("GetUint(https.port) = %v, want %v", got, uint(0))
		}
	})

	t.Run("GetUintOr", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUintOr("http.port", 3000); got != uint(8080) {
			t.Errorf("GetUintOr(http.port, 3000) = %v, want %v", got, uint(8080))
		}
		if got := GetUintOr("https.port", 8443); got != uint(8443) {
			t.Errorf("GetUintOr(https.port, 8443) = %v, want %v", got, uint(8443))
		}
	})

	t.Run("GetUint16", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint16("http.port"); got != uint16(8080) {
			t.Errorf("GetUint16(http.port) = %v, want %v", got, uint16(8080))
		}
		if got := GetUint16("https.port"); got != uint16(0) {
			t.Errorf("GetUint16(https.port) = %v, want %v", got, uint16(0))
		}
	})

	t.Run("GetUint16Or", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint16Or("http.port", 3000); got != uint16(8080) {
			t.Errorf("GetUint16Or(http.port, 3000) = %v, want %v", got, uint16(8080))
		}
		if got := GetUint16Or("https.port", 8443); got != uint16(8443) {
			t.Errorf("GetUint16Or(https.port, 8443) = %v, want %v", got, uint16(8443))
		}
	})

	t.Run("GetUint32", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint32("http.port"); got != uint32(8080) {
			t.Errorf("GetUint32(http.port) = %v, want %v", got, uint32(8080))
		}
		if got := GetUint32("https.port"); got != uint32(0) {
			t.Errorf("GetUint32(https.port) = %v, want %v", got, uint32(0))
		}
	})

	t.Run("GetUint32Or", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint32Or("http.port", 3000); got != uint32(8080) {
			t.Errorf("GetUint32Or(http.port, 3000) = %v, want %v", got, uint32(8080))
		}
		if got := GetUint32Or("https.port", 8443); got != uint32(8443) {
			t.Errorf("GetUint32Or(https.port, 8443) = %v, want %v", got, uint32(8443))
		}
	})

	t.Run("GetUint64", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint64("http.port"); got != uint64(8080) {
			t.Errorf("GetUint64(http.port) = %v, want %v", got, uint64(8080))
		}
		if got := GetUint64("https.port"); got != uint64(0) {
			t.Errorf("GetUint64(https.port) = %v, want %v", got, uint64(0))
		}
	})

	t.Run("GetUint64Or", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)

		if got := GetUint64Or("http.port", 3000); got != uint64(8080) {
			t.Errorf("GetUint64Or(http.port, 3000) = %v, want %v", got, uint64(8080))
		}
		if got := GetUint64Or("https.port", 8443); got != uint64(8443) {
			t.Errorf("GetUint64Or(https.port, 8443) = %v, want %v", got, uint64(8443))
		}
	})

	t.Run("GetFloat64", func(t *testing.T) {
		Reset()
		Set("stable_diffusion.cfg", 4.5)

		if got := GetFloat64("stable_diffusion.cfg"); got != 4.5 {
			t.Errorf("GetFloat64(stable_diffusion.cfg) = %v, want %v", got, 4.5)
		}
		if got := GetFloat64("stable_diffusion.prompt_strength"); got != float64(0) {
			t.Errorf("GetFloat64(stable_diffusion.prompt_strength) = %v, want %v", got, float64(0))
		}
	})

	t.Run("GetFloat64Or", func(t *testing.T) {
		Reset()
		Set("stable_diffusion.cfg", 4.5)

		if got := GetFloat64Or("stable_diffusion.cfg", 7); got != 4.5 {
			t.Errorf("GetFloat64Or(stable_diffusion.cfg, 7) = %v, want %v", got, 4.5)
		}
		if got := GetFloat64Or("stable_diffusion.prompt_strength", 0.85); got != 0.85 {
			t.Errorf("GetFloat64Or(stable_diffusion.prompt_strength, 0.85) = %v, want %v", got, 0.85)
		}
	})

	t.Run("GetDuration", func(t *testing.T) {
		Reset()
		Set("services.thrift.request_timeout_s", 60)

		if got := GetDuration("services.thrift.request_timeout_s"); got != time.Duration(60) {
			t.Errorf("GetDuration(services.thrift.request_timeout_s) = %v, want %v", got, time.Duration(60))
		}
		if got := GetDuration("services.thrift.request_timeout"); got != time.Duration(0) {
			t.Errorf("GetDuration(services.thrift.request_timeout) = %v, want %v", got, time.Duration(0))
		}
	})

	t.Run("GetDurationOr", func(t *testing.T) {
		Reset()
		Set("services.thrift.request_timeout_s", 60)

		if got := GetDurationOr("services.thrift.request_timeout_s", 30); got != time.Duration(60) {
			t.Errorf("GetDurationOr(services.thrift.request_timeout_s, 30) = %v, want %v", got, time.Duration(60))
		}
		if got := GetDurationOr("services.thrift.request_timeout", 60); got != time.Duration(60) {
			t.Errorf("GetDurationOr(services.thrift.request_timeout, 60) = %v, want %v", got, time.Duration(60))
		}
	})

	t.Run("GetStringPtr", func(t *testing.T) {
		Reset()
		Set("app.env", "local")

		result := GetStringPtr("app.env")
		if result == nil {
			t.Error("GetStringPtr(app.env) = nil, want non-nil")
		} else if *result != "local" {
			t.Errorf("*GetStringPtr(app.env) = %v, want %v", *result, "local")
		}

		result2 := GetStringPtr("app.environment")
		if result2 != nil {
			t.Errorf("GetStringPtr(app.environment) = %v, want nil", result2)
		}
	})

	t.Run("GetStringSlice", func(t *testing.T) {
		Reset()
		Set("app.hosts", []string{"localhost", "127.0.0.1"})

		got := GetStringSlice("app.hosts")
		want := []string{"localhost", "127.0.0.1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetStringSlice(app.hosts) = %v, want %v", got, want)
		}

		got2 := GetStringSlice("app.unknown")
		if got2 != nil {
			t.Errorf("GetStringSlice(app.unknown) = %v, want nil", got2)
		}
	})

	t.Run("GetIntSlice", func(t *testing.T) {
		Reset()
		Set("app.ports", []int{8080, 8081, 8082})

		got := GetIntSlice("app.ports")
		want := []int{8080, 8081, 8082}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetIntSlice(app.ports) = %v, want %v", got, want)
		}

		got2 := GetIntSlice("app.unknown")
		if got2 != nil {
			t.Errorf("GetIntSlice(app.unknown) = %v, want nil", got2)
		}
	})

	t.Run("GetStringMap", func(t *testing.T) {
		Reset()
		Set("app.metadata", map[string]interface{}{"key": "value", "count": 42})

		result := GetStringMap("app.metadata")
		if result["key"] != "value" {
			t.Errorf("GetStringMap(app.metadata)[key] = %v, want %v", result["key"], "value")
		}
		if result["count"] != 42 {
			t.Errorf("GetStringMap(app.metadata)[count] = %v, want %v", result["count"], 42)
		}

		got := GetStringMap("app.unknown")
		want := map[string]interface{}{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetStringMap(app.unknown) = %v, want %v", got, want)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		Reset()
		Set("app.env", "local")
		Set("app.debug", true)

		type AppConfig struct {
			App struct {
				Env   string `yaml:"env"`
				Debug bool   `yaml:"debug"`
			} `yaml:"app"`
		}

		var cfg AppConfig
		err := Unmarshal(&cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.App.Env != "local" {
			t.Errorf("cfg.App.Env = %v, want %v", cfg.App.Env, "local")
		}
		if cfg.App.Debug != true {
			t.Errorf("cfg.App.Debug = %v, want %v", cfg.App.Debug, true)
		}
	})

	t.Run("UnmarshalKey", func(t *testing.T) {
		Reset()
		Set("http.host", "localhost")
		Set("http.port", 8080)

		type HTTPConfig struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}

		var cfg HTTPConfig
		err := UnmarshalKey("http", &cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Host != "localhost" {
			t.Errorf("cfg.Host = %v, want %v", cfg.Host, "localhost")
		}
		if cfg.Port != 8080 {
			t.Errorf("cfg.Port = %v, want %v", cfg.Port, 8080)
		}
	})

	t.Run("AllSettings", func(t *testing.T) {
		Reset()
		Set("app.env", "local")
		Set("http.port", 8080)

		settings := AllSettings()
		if settings == nil {
			t.Error("AllSettings() = nil, want non-nil")
		}
		if _, ok := settings["app"]; !ok {
			t.Error("AllSettings() does not contain 'app' key")
		}
		if _, ok := settings["http"]; !ok {
			t.Error("AllSettings() does not contain 'http' key")
		}
	})
}

func TestGettersEnvOverride(t *testing.T) {
	t.Run("IsSet returns true for env var", func(t *testing.T) {
		Reset()
		os.Setenv("APP_NAME", "test")
		defer os.Unsetenv("APP_NAME")

		if !IsSet("app.name") {
			t.Error("IsSet(app.name) = false, want true when env var is set")
		}
	})

	t.Run("GetString prefers env var over config", func(t *testing.T) {
		Reset()
		Set("app.env", "from_config")
		os.Setenv("APP_ENV", "from_env")
		defer os.Unsetenv("APP_ENV")

		if got := GetString("app.env"); got != "from_env" {
			t.Errorf("GetString(app.env) = %q, want %q", got, "from_env")
		}
	})

	t.Run("GetBool prefers env var over config", func(t *testing.T) {
		Reset()
		Set("app.debug", false)
		os.Setenv("APP_DEBUG", "true")
		defer os.Unsetenv("APP_DEBUG")

		if got := GetBool("app.debug"); got != true {
			t.Errorf("GetBool(app.debug) = %v, want %v", got, true)
		}
	})

	t.Run("GetInt prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetInt("http.port"); got != 3000 {
			t.Errorf("GetInt(http.port) = %v, want %v", got, 3000)
		}
	})

	t.Run("GetInt32 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetInt32("http.port"); got != int32(3000) {
			t.Errorf("GetInt32(http.port) = %v, want %v", got, int32(3000))
		}
	})

	t.Run("GetInt64 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetInt64("http.port"); got != int64(3000) {
			t.Errorf("GetInt64(http.port) = %v, want %v", got, int64(3000))
		}
	})

	t.Run("GetUint prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetUint("http.port"); got != uint(3000) {
			t.Errorf("GetUint(http.port) = %v, want %v", got, uint(3000))
		}
	})

	t.Run("GetUint16 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetUint16("http.port"); got != uint16(3000) {
			t.Errorf("GetUint16(http.port) = %v, want %v", got, uint16(3000))
		}
	})

	t.Run("GetUint32 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetUint32("http.port"); got != uint32(3000) {
			t.Errorf("GetUint32(http.port) = %v, want %v", got, uint32(3000))
		}
	})

	t.Run("GetUint64 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("http.port", 8080)
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_PORT")

		if got := GetUint64("http.port"); got != uint64(3000) {
			t.Errorf("GetUint64(http.port) = %v, want %v", got, uint64(3000))
		}
	})

	t.Run("GetFloat64 prefers env var over config", func(t *testing.T) {
		Reset()
		Set("app.ratio", 1.5)
		os.Setenv("APP_RATIO", "2.5")
		defer os.Unsetenv("APP_RATIO")

		if got := GetFloat64("app.ratio"); got != 2.5 {
			t.Errorf("GetFloat64(app.ratio) = %v, want %v", got, 2.5)
		}
	})

	t.Run("GetDuration prefers env var over config", func(t *testing.T) {
		Reset()
		Set("app.timeout", 30)
		os.Setenv("APP_TIMEOUT", "60")
		defer os.Unsetenv("APP_TIMEOUT")

		if got := GetDuration("app.timeout"); got != time.Duration(60) {
			t.Errorf("GetDuration(app.timeout) = %v, want %v", got, time.Duration(60))
		}
	})

	t.Run("GetDuration parses duration string from env var", func(t *testing.T) {
		Reset()
		os.Setenv("APP_TIMEOUT", "30s")
		defer os.Unsetenv("APP_TIMEOUT")

		if got := GetDuration("app.timeout"); got != 30*time.Second {
			t.Errorf("GetDuration(app.timeout) = %v, want %v", got, 30*time.Second)
		}
	})

	t.Run("AllSettings applies env var overrides", func(t *testing.T) {
		Reset()
		Set("app.env", "from_config")
		Set("app.debug", true)
		Set("http.port", 8080)

		os.Setenv("APP_ENV", "from_env")
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("APP_ENV")
		defer os.Unsetenv("HTTP_PORT")

		settings := AllSettings()

		// Check that env var overrides are applied
		appSettings, ok := settings["app"].(map[string]any)
		if !ok {
			t.Fatal("AllSettings()[app] is not a map")
		}
		if appSettings["env"] != "from_env" {
			t.Errorf("AllSettings()[app][env] = %v, want %v", appSettings["env"], "from_env")
		}

		httpSettings, ok := settings["http"].(map[string]any)
		if !ok {
			t.Fatal("AllSettings()[http] is not a map")
		}
		// Env var is converted to match original type (int)
		if httpSettings["port"] != 3000 {
			t.Errorf("AllSettings()[http][port] = %v, want %v", httpSettings["port"], 3000)
		}
	})

	t.Run("Unmarshal applies env var overrides", func(t *testing.T) {
		Reset()
		Set("app.env", "from_config")
		Set("app.debug", false)

		os.Setenv("APP_ENV", "from_env")
		os.Setenv("APP_DEBUG", "true")
		defer os.Unsetenv("APP_ENV")
		defer os.Unsetenv("APP_DEBUG")

		type Config struct {
			App struct {
				Env   string `yaml:"env"`
				Debug bool   `yaml:"debug"`
			} `yaml:"app"`
		}

		var cfg Config
		if err := Unmarshal(&cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.App.Env != "from_env" {
			t.Errorf("cfg.App.Env = %v, want %v", cfg.App.Env, "from_env")
		}
		if cfg.App.Debug != true {
			t.Errorf("cfg.App.Debug = %v, want %v", cfg.App.Debug, true)
		}
	})

	t.Run("UnmarshalKey applies env var overrides", func(t *testing.T) {
		Reset()
		Set("http.host", "localhost")
		Set("http.port", 8080)

		os.Setenv("HTTP_HOST", "0.0.0.0")
		os.Setenv("HTTP_PORT", "3000")
		defer os.Unsetenv("HTTP_HOST")
		defer os.Unsetenv("HTTP_PORT")

		type HTTPConfig struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}

		var cfg HTTPConfig
		if err := UnmarshalKey("http", &cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Host != "0.0.0.0" {
			t.Errorf("cfg.Host = %v, want %v", cfg.Host, "0.0.0.0")
		}
		if cfg.Port != 3000 {
			t.Errorf("cfg.Port = %v, want %v", cfg.Port, 3000)
		}
	})

	t.Run("GetStringSlice prefers env var over config", func(t *testing.T) {
		Reset()
		Set("cors.allowed_origins", []string{"https://config.com"})

		os.Setenv("CORS_ALLOWED_ORIGINS", "https://env1.com,https://env2.com")
		defer os.Unsetenv("CORS_ALLOWED_ORIGINS")

		got := GetStringSlice("cors.allowed_origins")
		want := []string{"https://env1.com", "https://env2.com"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetStringSlice(cors.allowed_origins) = %v, want %v", got, want)
		}
	})

	t.Run("GetStringSlice env var with whitespace", func(t *testing.T) {
		Reset()
		os.Setenv("TAGS", "foo, bar, baz")
		defer os.Unsetenv("TAGS")

		got := GetStringSlice("tags")
		want := []string{"foo", "bar", "baz"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetStringSlice(tags) = %v, want %v", got, want)
		}
	})

	t.Run("GetStringSlice env var filters empty values", func(t *testing.T) {
		Reset()
		os.Setenv("ITEMS", "a,,b,,c")
		defer os.Unsetenv("ITEMS")

		got := GetStringSlice("items")
		want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetStringSlice(items) = %v, want %v", got, want)
		}
	})

	t.Run("GetIntSlice prefers env var over config", func(t *testing.T) {
		Reset()
		Set("retry.backoff_ms", []int{100, 200})

		os.Setenv("RETRY_BACKOFF_MS", "500,1000,2000")
		defer os.Unsetenv("RETRY_BACKOFF_MS")

		got := GetIntSlice("retry.backoff_ms")
		want := []int{500, 1000, 2000}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetIntSlice(retry.backoff_ms) = %v, want %v", got, want)
		}
	})

	t.Run("GetIntSlice env var with whitespace", func(t *testing.T) {
		Reset()
		os.Setenv("PORTS", "8080, 8081, 8082")
		defer os.Unsetenv("PORTS")

		got := GetIntSlice("ports")
		want := []int{8080, 8081, 8082}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetIntSlice(ports) = %v, want %v", got, want)
		}
	})

	t.Run("GetIntSlice env var filters invalid values", func(t *testing.T) {
		Reset()
		os.Setenv("NUMBERS", "1,invalid,2,,3")
		defer os.Unsetenv("NUMBERS")

		got := GetIntSlice("numbers")
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetIntSlice(numbers) = %v, want %v", got, want)
		}
	})

	t.Run("AllSettings applies env var overrides for slices", func(t *testing.T) {
		Reset()
		Set("app.tags", []string{"from", "config"})
		Set("app.ports", []int{8080})

		os.Setenv("APP_TAGS", "from,env")
		os.Setenv("APP_PORTS", "3000,3001")
		defer os.Unsetenv("APP_TAGS")
		defer os.Unsetenv("APP_PORTS")

		settings := AllSettings()

		appSettings, ok := settings["app"].(map[string]any)
		if !ok {
			t.Fatal("AllSettings()[app] is not a map")
		}

		gotTags, ok := appSettings["tags"].([]string)
		if !ok {
			t.Fatalf("AllSettings()[app][tags] is not []string, got %T", appSettings["tags"])
		}
		wantTags := []string{"from", "env"}
		if !reflect.DeepEqual(gotTags, wantTags) {
			t.Errorf("AllSettings()[app][tags] = %v, want %v", gotTags, wantTags)
		}

		gotPorts, ok := appSettings["ports"].([]int)
		if !ok {
			t.Fatalf("AllSettings()[app][ports] is not []int, got %T", appSettings["ports"])
		}
		wantPorts := []int{3000, 3001}
		if !reflect.DeepEqual(gotPorts, wantPorts) {
			t.Errorf("AllSettings()[app][ports] = %v, want %v", gotPorts, wantPorts)
		}
	})

	t.Run("Unmarshal applies env var overrides for slices", func(t *testing.T) {
		Reset()
		Set("app.tags", []string{"from", "config"})

		os.Setenv("APP_TAGS", "tag1,tag2,tag3")
		defer os.Unsetenv("APP_TAGS")

		type Config struct {
			App struct {
				Tags []string `yaml:"tags"`
			} `yaml:"app"`
		}

		var cfg Config
		if err := Unmarshal(&cfg); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := []string{"tag1", "tag2", "tag3"}
		if !reflect.DeepEqual(cfg.App.Tags, want) {
			t.Errorf("cfg.App.Tags = %v, want %v", cfg.App.Tags, want)
		}
	})
}

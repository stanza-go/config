# config

A strongly opinionated configuration package for Go services.

## Philosophy

When starting a new HTTP/gRPC service, I prefer to spend zero time thinking about configuration — how to load it, which
format to use, or how environment variables should work. I just want it done, the same way, every time.

Go gives developers freedom to do things in many ways. That's great for libraries and frameworks, but for everyday
services, that freedom can become decision fatigue. This package trades flexibility for simplicity: **fewer choices,
more focus on business logic**.

Core principles:

- **Convention over configuration** — sensible defaults, minimal setup
- **Fail fast** — broken config crashes immediately at startup
- **Environment variables always win** — easy deployment across environments
- **One way to do things** — no multiple approaches for the same problem
- **Minimal dependencies** — only `gopkg.in/yaml.v3`, fewer CVEs to worry about

## Installation

```bash
go get github.com/stanza-go/config
```

## Quick Start

### 1. Create `config.yaml`

```yaml
# config.yaml

app:
  env: local
  service_name: my-service
  debug: false

http:
  host: localhost
  port: 8080
```

### 2. Initialize and use

```go
package main

import (
	"fmt"

	"github.com/stanza-go/config"
)

func main() {
	config.Init()

	env := config.GetString("app.env")
	port := config.GetInt("http.port")
	debug := config.GetBool("app.debug")

	fmt.Printf("Running %s on port %d (debug: %v)\n", env, port, debug)
}
```

### 3. Override with environment variables

```bash
APP_ENV=production HTTP_PORT=3000 go run main.go
```

Or use a `.env` file (loaded automatically if present):

```
# .env

APP_ENV=production
HTTP_PORT=3000
```

## Opinions

This package enforces the following conventions. If they don't fit your use case, this package may not be for you.

| Opinion                                           | Reasoning                                   |
|---------------------------------------------------|---------------------------------------------|
| Config file must be named `config.yaml`           | One name, no confusion                      |
| Environment variables always override config file | Easy deployment across environments         |
| Use `snake_case` for YAML keys                    | Maps cleanly to `UPPER_SNAKE_CASE` env vars |
| Dot notation for nested keys                      | `app.service_name` → `APP_SERVICE_NAME`     |
| Fatal on startup if config is broken              | Fail fast, fix before deploying             |
| Single config file, no merging                    | Simplicity over flexibility                 |

### Naming Convention

Use `snake_case` for YAML keys — it maps cleanly to environment variables:

| YAML Key           | Environment Variable | Status                   |
|--------------------|----------------------|--------------------------|
| `app.service_name` | `APP_SERVICE_NAME`   | ✅ Recommended            |
| `app.serviceName`  | `APP_SERVICENAME`    | ⚠️ Confusing             |
| `app.service.name` | `APP_SERVICE_NAME`   | ❌ Conflicts with nesting |

## Environment Variable Override

Every getter checks environment variables first. The key is converted from dot notation to `UPPER_SNAKE_CASE`:

```
app.service_name  →  APP_SERVICE_NAME
http.port         →  HTTP_PORT
database.max_conn →  DATABASE_MAX_CONN
```

### Slice Values (Comma-Separated)

For slices, use comma-separated values in environment variables:

```bash
# String slice
export CORS_ALLOWED_ORIGINS=https://a.com,https://b.com

# Int slice
export RETRY_BACKOFF_MS=100,200,500,1000
```

Whitespace around commas is trimmed, empty values are filtered:

```bash
export TAGS="foo, bar, baz"     # ["foo", "bar", "baz"]
export PORTS="8080,,8081"       # [8080, 8081]
```

## Available Functions

### Getters

| Function                      | Return Type      | Env Override        |
|-------------------------------|------------------|---------------------|
| `GetString(key)`              | `string`         | ✅                   |
| `GetStringOr(key, default)`   | `string`         | ✅                   |
| `GetStringPtr(key)`           | `*string`        | ✅                   |
| `GetBool(key)`                | `bool`           | ✅                   |
| `GetBoolOr(key, default)`     | `bool`           | ✅                   |
| `GetInt(key)`                 | `int`            | ✅                   |
| `GetIntOr(key, default)`      | `int`            | ✅                   |
| `GetInt32(key)`               | `int32`          | ✅                   |
| `GetInt32Or(key, default)`    | `int32`          | ✅                   |
| `GetInt64(key)`               | `int64`          | ✅                   |
| `GetInt64Or(key, default)`    | `int64`          | ✅                   |
| `GetUint(key)`                | `uint`           | ✅                   |
| `GetUintOr(key, default)`     | `uint`           | ✅                   |
| `GetUint16(key)`              | `uint16`         | ✅                   |
| `GetUint16Or(key, default)`   | `uint16`         | ✅                   |
| `GetUint32(key)`              | `uint32`         | ✅                   |
| `GetUint32Or(key, default)`   | `uint32`         | ✅                   |
| `GetUint64(key)`              | `uint64`         | ✅                   |
| `GetUint64Or(key, default)`   | `uint64`         | ✅                   |
| `GetFloat64(key)`             | `float64`        | ✅                   |
| `GetFloat64Or(key, default)`  | `float64`        | ✅                   |
| `GetDuration(key)`            | `time.Duration`  | ✅                   |
| `GetDurationOr(key, default)` | `time.Duration`  | ✅                   |
| `GetStringSlice(key)`         | `[]string`       | ✅ (comma-separated) |
| `GetIntSlice(key)`            | `[]int`          | ✅ (comma-separated) |
| `GetStringMap(key)`           | `map[string]any` | ❌                   |
| `IsSet(key)`                  | `bool`           | ✅                   |

### Unmarshaling

| Function               | Description                            | Env Override |
|------------------------|----------------------------------------|--------------|
| `Unmarshal(v)`         | Unmarshal entire config into struct    | ✅            |
| `UnmarshalKey(key, v)` | Unmarshal specific section into struct | ✅            |
| `AllSettings()`        | Get all settings as map                | ✅            |

### Testing Utilities

These functions are intended for testing only:

| Function          | Description                         |
|-------------------|-------------------------------------|
| `Set(key, value)` | Set a config value programmatically |
| `Reset()`         | Clear all config data               |

```go
func TestSomething(t *testing.T) {
    config.Reset()
    config.Set("app.env", "test")
    // ... run test
}
```

## How It Works

1. **Config file discovery**: `Init()` looks for `config.yaml` starting from the current working directory, traversing
   up to parent directories until found.

2. **Environment file loading**: If a `.env` file exists in the same directory as `config.yaml`, it is loaded
   automatically.

3. **Value retrieval**: Every getter checks environment variables first (converted to `UPPER_SNAKE_CASE`), then falls
   back to the config file value.

## When NOT to Use This Package

This package is designed for HTTP/gRPC services with simple configuration needs. It's not suitable for:

- ❌ CLI tools requiring flags/arguments — consider [Cobra](https://github.com/spf13/cobra)
- ❌ Apps needing multiple config formats (JSON, TOML, HCL)
- ❌ Apps requiring hot-reload / watch config changes
- ❌ Apps needing config validation at startup
- ❌ Apps requiring complex config merging from multiple sources

For those use cases, [Viper](https://github.com/spf13/viper) or [Koanf](https://github.com/knadh/koanf) are excellent
alternatives.

## Example

> Example project coming soon.

## License

MIT License

## Issues

Issues and feedback are welcome.

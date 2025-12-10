package config

import (
	"strconv"
	"strings"
	"time"
)

// Type conversion helpers for converting values from YAML or environment variables
// to specific Go types.

func toString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return ""
	}
}

func toBool(v any) bool {
	switch val := v.(type) {
	case bool:
		return val
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val != 0
	case int64:
		return val != 0
	case float64:
		return val != 0
	default:
		return false
	}
}

func toInt(v any) int {
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	case string:
		i, _ := strconv.Atoi(val)
		return i
	default:
		return 0
	}
}

func toInt32(v any) int32 {
	switch val := v.(type) {
	case int:
		return int32(val)
	case int32:
		return val
	case int64:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		i, _ := strconv.ParseInt(val, 10, 32)
		return int32(i)
	default:
		return 0
	}
}

func toInt64(v any) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case float64:
		return int64(val)
	case string:
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	default:
		return 0
	}
}

func toUint(v any) uint {
	switch val := v.(type) {
	case int:
		return uint(val)
	case int64:
		return uint(val)
	case uint:
		return val
	case uint64:
		return uint(val)
	case float64:
		return uint(val)
	case string:
		i, _ := strconv.ParseUint(val, 10, 64)
		return uint(i)
	default:
		return 0
	}
}

func toUint16(v any) uint16 {
	switch val := v.(type) {
	case int:
		return uint16(val)
	case int64:
		return uint16(val)
	case uint:
		return uint16(val)
	case uint16:
		return val
	case uint64:
		return uint16(val)
	case float64:
		return uint16(val)
	case string:
		i, _ := strconv.ParseUint(val, 10, 16)
		return uint16(i)
	default:
		return 0
	}
}

func toUint32(v any) uint32 {
	switch val := v.(type) {
	case int:
		return uint32(val)
	case int64:
		return uint32(val)
	case uint:
		return uint32(val)
	case uint32:
		return val
	case uint64:
		return uint32(val)
	case float64:
		return uint32(val)
	case string:
		i, _ := strconv.ParseUint(val, 10, 32)
		return uint32(i)
	default:
		return 0
	}
}

func toUint64(v any) uint64 {
	switch val := v.(type) {
	case int:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return val
	case float64:
		return uint64(val)
	case string:
		i, _ := strconv.ParseUint(val, 10, 64)
		return i
	default:
		return 0
	}
}

func toFloat64(v any) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}

func toDuration(v any) time.Duration {
	switch val := v.(type) {
	case time.Duration:
		return val
	case int:
		return time.Duration(val)
	case int64:
		return time.Duration(val)
	case float64:
		return time.Duration(val)
	case string:
		// Try parsing as duration string first (e.g., "30s", "1h")
		if d, err := time.ParseDuration(val); err == nil {
			return d
		}
		// Fall back to parsing as integer
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return time.Duration(i)
		}
		return 0
	default:
		return 0
	}
}

func toStringSlice(v any) []string {
	switch val := v.(type) {
	case []string:
		return val
	case []any:
		result := make([]string, len(val))
		for i, item := range val {
			result[i] = toString(item)
		}
		return result
	default:
		return nil
	}
}

func toIntSlice(v any) []int {
	switch val := v.(type) {
	case []int:
		return val
	case []any:
		result := make([]int, len(val))
		for i, item := range val {
			result[i] = toInt(item)
		}
		return result
	default:
		return nil
	}
}

func toStringMap(v any) map[string]any {
	switch val := v.(type) {
	case map[string]any:
		return val
	default:
		return map[string]any{}
	}
}

// splitAndTrimStringSlice splits a comma-separated string into a slice of strings.
// It trims whitespace from each element and filters out empty strings.
//
// Examples:
//
//	"a,b,c"       -> ["a", "b", "c"]
//	"a, b, c"     -> ["a", "b", "c"]
//	"a,,b"        -> ["a", "b"]
//	""            -> []
//	"  "          -> []
func splitAndTrimStringSlice(s string) []string {
	if strings.TrimSpace(s) == "" {
		return []string{}
	}

	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// splitAndTrimIntSlice splits a comma-separated string into a slice of integers.
// It trims whitespace from each element, filters out empty strings,
// and skips values that cannot be parsed as integers.
//
// Examples:
//
//	"1,2,3"       -> [1, 2, 3]
//	"1, 2, 3"     -> [1, 2, 3]
//	"1,,2"        -> [1, 2]
//	"1,invalid,2" -> [1, 2]
//	""            -> []
func splitAndTrimIntSlice(s string) []int {
	if strings.TrimSpace(s) == "" {
		return []int{}
	}

	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		if i, err := strconv.Atoi(trimmed); err == nil {
			result = append(result, i)
		}
	}

	return result
}

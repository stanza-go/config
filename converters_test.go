package config

import (
	"reflect"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int64", int64(100), "100"},
		{"float64", 3.14, "3.14"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"nil", nil, ""},
		{"unsupported type", []int{1, 2, 3}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.input); got != tt.want {
				t.Errorf("toString(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{"bool true", true, true},
		{"bool false", false, false},
		{"string true", "true", true},
		{"string false", "false", false},
		{"string 1", "1", true},
		{"string 0", "0", false},
		{"int non-zero", 1, true},
		{"int zero", 0, false},
		{"int64 non-zero", int64(100), true},
		{"int64 zero", int64(0), false},
		{"float64 non-zero", 1.5, true},
		{"float64 zero", 0.0, false},
		{"nil", nil, false},
		{"unsupported type", []int{1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBool(tt.input); got != tt.want {
				t.Errorf("toBool(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  int
	}{
		{"int", 42, 42},
		{"int64", int64(100), 100},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt(tt.input); got != tt.want {
				t.Errorf("toInt(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  int32
	}{
		{"int", 42, 42},
		{"int32", int32(100), 100},
		{"int64", int64(200), 200},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt32(tt.input); got != tt.want {
				t.Errorf("toInt32(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  int64
	}{
		{"int", 42, 42},
		{"int32", int32(100), 100},
		{"int64", int64(200), 200},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt64(tt.input); got != tt.want {
				t.Errorf("toInt64(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  uint
	}{
		{"int", 42, 42},
		{"int64", int64(100), 100},
		{"uint", uint(200), 200},
		{"uint64", uint64(300), 300},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUint(tt.input); got != tt.want {
				t.Errorf("toUint(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  uint16
	}{
		{"int", 42, 42},
		{"int64", int64(100), 100},
		{"uint", uint(200), 200},
		{"uint16", uint16(300), 300},
		{"uint64", uint64(400), 400},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUint16(tt.input); got != tt.want {
				t.Errorf("toUint16(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  uint32
	}{
		{"int", 42, 42},
		{"int64", int64(100), 100},
		{"uint", uint(200), 200},
		{"uint32", uint32(300), 300},
		{"uint64", uint64(400), 400},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUint32(tt.input); got != tt.want {
				t.Errorf("toUint32(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  uint64
	}{
		{"int", 42, 42},
		{"int64", int64(100), 100},
		{"uint", uint(200), 200},
		{"uint32", uint32(300), 300},
		{"uint64", uint64(400), 400},
		{"float64", 3.9, 3},
		{"string", "123", 123},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUint64(tt.input); got != tt.want {
				t.Errorf("toUint64(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  float64
	}{
		{"float64", 3.14, 3.14},
		{"float32", float32(2.5), 2.5},
		{"int", 42, 42.0},
		{"int64", int64(100), 100.0},
		{"string", "3.14", 3.14},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toFloat64(tt.input)
			// Use tolerance for float comparison
			if got != tt.want && (got-tt.want > 0.0001 || tt.want-got > 0.0001) {
				t.Errorf("toFloat64(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToDuration(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  time.Duration
	}{
		{"duration", time.Second * 30, time.Second * 30},
		{"int", 1000, time.Duration(1000)},
		{"int64", int64(2000), time.Duration(2000)},
		{"float64", 3000.0, time.Duration(3000)},
		{"string duration", "30s", time.Second * 30},
		{"string duration hours", "1h", time.Hour},
		{"string integer", "1000", time.Duration(1000)},
		{"string invalid", "abc", 0},
		{"nil", nil, 0},
		{"unsupported type", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDuration(tt.input); got != tt.want {
				t.Errorf("toDuration(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  []string
	}{
		{"string slice", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"any slice with strings", []any{"x", "y", "z"}, []string{"x", "y", "z"}},
		{"any slice with mixed types", []any{"hello", 42, true}, []string{"hello", "42", "true"}},
		{"nil", nil, nil},
		{"unsupported type", "not a slice", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringSlice(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringSlice(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToIntSlice(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  []int
	}{
		{"int slice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"any slice with ints", []any{10, 20, 30}, []int{10, 20, 30}},
		{"any slice with mixed numeric", []any{1, int64(2), 3.0}, []int{1, 2, 3}},
		{"nil", nil, nil},
		{"unsupported type", "not a slice", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toIntSlice(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toIntSlice(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToStringMap(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  map[string]any
	}{
		{"string map", map[string]any{"key": "value"}, map[string]any{"key": "value"}},
		{"nil", nil, map[string]any{}},
		{"unsupported type", "not a map", map[string]any{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringMap(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringMap(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSplitAndTrimStringSlice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "simple comma separated",
			input: "a,b,c",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "with whitespace around values",
			input: "a, b, c",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "with extra whitespace",
			input: "  a  ,  b  ,  c  ",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "with empty values filtered out",
			input: "a,,b",
			want:  []string{"a", "b"},
		},
		{
			name:  "with multiple empty values",
			input: "a,,,b,,c",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "trailing comma",
			input: "a,b,c,",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "leading comma",
			input: ",a,b,c",
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "empty string",
			input: "",
			want:  []string{},
		},
		{
			name:  "whitespace only",
			input: "   ",
			want:  []string{},
		},
		{
			name:  "single value",
			input: "single",
			want:  []string{"single"},
		},
		{
			name:  "single value with whitespace",
			input: "  single  ",
			want:  []string{"single"},
		},
		{
			name:  "urls",
			input: "https://example.com,https://api.example.com",
			want:  []string{"https://example.com", "https://api.example.com"},
		},
		{
			name:  "urls with spaces",
			input: "https://example.com, https://api.example.com",
			want:  []string{"https://example.com", "https://api.example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitAndTrimStringSlice(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitAndTrimStringSlice(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSplitAndTrimIntSlice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "simple comma separated",
			input: "1,2,3",
			want:  []int{1, 2, 3},
		},
		{
			name:  "with whitespace around values",
			input: "1, 2, 3",
			want:  []int{1, 2, 3},
		},
		{
			name:  "with extra whitespace",
			input: "  1  ,  2  ,  3  ",
			want:  []int{1, 2, 3},
		},
		{
			name:  "with empty values filtered out",
			input: "1,,2",
			want:  []int{1, 2},
		},
		{
			name:  "with multiple empty values",
			input: "1,,,2,,3",
			want:  []int{1, 2, 3},
		},
		{
			name:  "with invalid values filtered out",
			input: "1,invalid,2",
			want:  []int{1, 2},
		},
		{
			name:  "with mixed invalid values",
			input: "1,abc,2,def,3",
			want:  []int{1, 2, 3},
		},
		{
			name:  "trailing comma",
			input: "1,2,3,",
			want:  []int{1, 2, 3},
		},
		{
			name:  "leading comma",
			input: ",1,2,3",
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty string",
			input: "",
			want:  []int{},
		},
		{
			name:  "whitespace only",
			input: "   ",
			want:  []int{},
		},
		{
			name:  "single value",
			input: "42",
			want:  []int{42},
		},
		{
			name:  "single value with whitespace",
			input: "  42  ",
			want:  []int{42},
		},
		{
			name:  "negative numbers",
			input: "-1,-2,-3",
			want:  []int{-1, -2, -3},
		},
		{
			name:  "mixed positive and negative",
			input: "-1,0,1",
			want:  []int{-1, 0, 1},
		},
		{
			name:  "port numbers",
			input: "8080,8081,8082",
			want:  []int{8080, 8081, 8082},
		},
		{
			name:  "all invalid returns empty",
			input: "abc,def,ghi",
			want:  []int{},
		},
		{
			name:  "float values not parsed",
			input: "1.5,2.9,3.1",
			want:  []int{}, // strconv.Atoi doesn't parse floats
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitAndTrimIntSlice(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitAndTrimIntSlice(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

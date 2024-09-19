package stacktrace

import (
	"testing"
)

func TestFuncname(t *testing.T) {
	tests := []struct {
		Str  string
		Want string
	}{
		{"package", "package"},
		{"path/package", "package"},
		{"path/path/package", "package"},

		{"package.function", "function"},
		{"path/package.function", "function"},
		{"path/path/package.function", "function"},

		{"package.function[...]", "function[...]"},
		{"path/package.function[...]", "function[...]"},
		{"path/path/package.function[...]", "function[...]"},
	}
	for i, tt := range tests {
		got := funcname(tt.Str)
		if got != tt.Want {
			t.Errorf("%d got:%q want:%q", i, got, tt.Want)
		}
	}
}

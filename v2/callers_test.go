package stacktrace_test

import (
	"runtime"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestCallers(t *testing.T) {
	var f func(int) []uintptr
	f = func(n int) []uintptr {
		if n <= 0 {
			return stacktrace.Callers(0)
		}
		return f(n - 1)
	}
	callers := f(32)
	frames := runtime.CallersFrames(callers)
	for {
		frame, more := frames.Next()
		t.Log(frame.Function)
		if !more {
			break
		}
	}
}

func TestCallersLimit(t *testing.T) {
	testCases := []struct {
		name      string
		skip      int
		limit     int
		expectLen int
	}{
		{name: "Skip 0, Limit 0", skip: 0, limit: 0, expectLen: 0},
		{name: "Skip 0, Limit 1", skip: 0, limit: 1, expectLen: 1},
		{name: "Skip 0, Limit 2", skip: 0, limit: 2, expectLen: 2},
		{name: "Skip 1, Limit 1", skip: 1, limit: 1, expectLen: 1},
		{name: "Skip 1, Limit 2", skip: 1, limit: 2, expectLen: 2},
		{name: "Skip 2, Limit 1", skip: 2, limit: 1, expectLen: 1},
		{name: "Skip 0, Limit 99", skip: 0, limit: 99, expectLen: -1},
		{name: "Skip 1, Limit 99", skip: 1, limit: 99, expectLen: -1},
		{name: "Skip 2, Limit 99", skip: 2, limit: 99, expectLen: -1},
		{name: "Skip 0, Limit -1", skip: 0, limit: -1, expectLen: -1}, // Should return all frames after skip + 1
		{name: "Skip 1, Limit -1", skip: 1, limit: -1, expectLen: -1}, // Should return all frames after skip + 1
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectLen == -1 {
				pcs := make([]uintptr, 32)
				n := runtime.Callers(tc.skip+1, pcs)
				tc.expectLen = n
			}
			result := stacktrace.CallersLimit(tc.skip, tc.limit)
			if len(result) != tc.expectLen {
				t.Errorf("Expected length %d, got %d", tc.expectLen, len(result))
			}
		})
	}
}

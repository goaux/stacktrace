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

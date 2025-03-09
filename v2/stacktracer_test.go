package stacktrace_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

type testTracer struct {
	callers []uintptr
}

var _ stacktrace.StackTracer = (*testTracer)(nil)

func newTestTracer() error {
	return &testTracer{callers: stacktrace.Callers(0)}
}

func newTestTracer2() (any, error) {
	return nil, newTestTracer()
}

func newTestTracer3() (any, any, error) {
	return nil, nil, newTestTracer()
}

func newTestTracer4() (any, any, any, error) {
	return nil, nil, nil, newTestTracer()
}

func (tt *testTracer) Error() string {
	return "testTracer"
}

func (tt *testTracer) StackTrace() []uintptr {
	return tt.callers
}

func TestStackTracer(t *testing.T) {
	check := func(t *testing.T, err error) {
		t.Helper()
		info := stacktrace.GetDebugInfo(err)
		if len(info.StackEntries) == 0 {
			t.Errorf("err must have StackEntries")
		}
		want := "/stacktracer_test.go:18 newTestTracer\n"
		got := info.Format()
		if !strings.Contains(got, want) {
			t.Errorf("got=%q", got)
		}
	}

	t.Run("Errorf", func(t *testing.T) {
		err := fmt.Errorf("error: %w", newTestTracer())
		check(t, err)
	})

	t.Run("Trace", func(t *testing.T) {
		err := stacktrace.Trace(newTestTracer())
		check(t, err)
	})

	t.Run("Trace2", func(t *testing.T) {
		_, err := stacktrace.Trace2(newTestTracer2())
		check(t, err)
	})

	t.Run("Trace3", func(t *testing.T) {
		_, _, err := stacktrace.Trace3(newTestTracer3())
		check(t, err)
	})

	t.Run("Trace4", func(t *testing.T) {
		_, _, _, err := stacktrace.Trace4(newTestTracer4())
		check(t, err)
	})
}

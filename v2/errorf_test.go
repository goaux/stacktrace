package stacktrace_test

import (
	"os"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestErrorf(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		err := stacktrace.Errorf("42 %w", os.ErrInvalid)
		info := stacktrace.GetDebugInfo(err)
		want := "42 invalid argument (errorf_test.go:12 TestErrorf.func1)"
		if got := info.Detail; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
		if n := len(info.StackEntries); n != 3 {
			t.Errorf("len(info.StackEntries) = %d, must be 3", n)
		}
	})
	t.Run("single", func(t *testing.T) {
		err := stacktrace.Errorf("42 %w", os.ErrInvalid)
		err = stacktrace.Errorf("41 %w", err)
		err = stacktrace.Errorf("40 %w", err)
		info := stacktrace.GetDebugInfo(err)
		want := "40 41 42 invalid argument (errorf_test.go:23 TestErrorf.func2)"
		if got := info.Detail; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
		if n := len(info.StackEntries); n != 4 {
			t.Logf("\n%s", info.Format())
			t.Errorf("len(info.StackEntries) = %d, must be 4", n)
		}
	})
}

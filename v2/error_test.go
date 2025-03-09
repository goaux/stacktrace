package stacktrace_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestError(t *testing.T) {
	t.Run("zerovalue", func(t *testing.T) {
		want := ""
		err := new(stacktrace.Error)
		got := err.Error()
		if got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
	})

	t.Run("unwrap", func(t *testing.T) {
		err := stacktrace.Trace(fmt.Errorf("%w", os.ErrInvalid))
		if !errors.Is(err, os.ErrInvalid) {
			t.Error("err must be os.ErrInvalid")
		}
	})

	t.Run("stacktracer", func(t *testing.T) {
		var err error = stacktrace.NewError(nil, []uintptr{1, 2, 3})
		if v, ok := err.(stacktrace.StackTracer); ok {
			got := v.StackTrace()
			want := []uintptr{1, 2, 3}
			if !equal(got, want) {
				t.Errorf("got=%v want=%v", got, want)
			}
		}
	})
}

// equal returns true if `s1` and `s2` are identical.
//
// In Go 1.20, [slices.Equal] was not available, so this custom implementation
// was created. Once support for Go versions 1.21 or later is ensured, it is
// recommended to use [slices.Equal` and remove this function.
func equal(s1, s2 []uintptr) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

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
}

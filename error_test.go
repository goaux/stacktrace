package stacktrace_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/goaux/stacktrace"
	v2 "github.com/goaux/stacktrace/v2"
)

func TestError(t *testing.T) {
	err0 := errors.New("err0")
	err1 := fmt.Errorf("%w", err0)
	err := &stacktrace.Error{Cause: err1}

	assert.Equal(t, err0.Error(), err.Error())

	assert.Equal(t, err1, errors.Unwrap(err))

	assert.True(t, errors.Is(err, err0))
	assert.True(t, errors.Is(err, err1))
}

func TestExtract(t *testing.T) {
	t.Run("v2", func(t *testing.T) {
		// stacktrace(v1) knows abount v2.Error,
		t.Run("format", func(t *testing.T) {
			lines := strings.Split(stacktrace.Format(v2.New("v2")), "\n")
			if len(lines) <= 1 {
				t.Error("len(lines) must be greater than 1")
			}
		})
		t.Run("extract", func(t *testing.T) {
			n := stacktrace.Extract(v2.New("v2"))
			if len(n) != 1 {
				t.Errorf("len(n) must be 1, got=%d", len(n))
			}
		})
	})
}

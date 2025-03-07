package stacktrace_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/goaux/stacktrace"
)

func TestError(t *testing.T) {
	err0 := errors.New("err0")
	err1 := fmt.Errorf("%w", err0)
	err := &stacktrace.Error{Cause: err1}

	assert.Equal(t, err0.Error(), err.Error())

	assert.Equal(t, err1, errors.Unwrap(err))

	assert.True(t, errors.Is(err, err0))
	assert.True(t, errors.Is(err, err1))

	t.Run("stacktrace", func(t *testing.T) {
		var err error = &stacktrace.Error{Frames: []uintptr{1, 2, 3}}
		if v, ok := err.(stacktrace.StackTracer); ok {
			got := v.StackTrace()
			want := []uintptr{1, 2, 3}
			if !slices.Equal(got, want) {
				t.Errorf("got=%v want=%v", got, want)
			}
		}
	})
}

func TestWith(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := stacktrace.With(nil)
		if got != nil {
			t.Errorf("must be nil, got=%v", got)
		}
	})
}

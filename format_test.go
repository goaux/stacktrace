package stacktrace_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/goaux/stacktrace"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		err := hello()
		err = stacktrace.With(fmt.Errorf("this is %w", err))
		err = stacktrace.With(err)

		assert.Equal(t, `this is ERROR`, err.Error())

		assert.Contains(t,
			stacktrace.Format(err),
			"this is ERROR\n  [0]: ERROR\n\t",
		)

		assert.Equal(t, `this is ERROR`, fmt.Sprintf("%v", err))
		assert.Equal(t, `this is ERROR`, fmt.Sprintf("%+v", err))

		assert.Contains(t,
			fmt.Sprintf("%#v", err),
			`&fmt.wrapError{msg:"this is ERROR", err:(*stacktrace.Error)`,
		)

		assert.Len(t, stacktrace.Dump(err).Traces, 1)

		t.Log(stacktrace.Format(err))
	})

	t.Run("always", func(t *testing.T) {
		always := stacktrace.Always
		err := hello(always, stacktrace.Limit(2))
		err = stacktrace.With(fmt.Errorf("this is %w", err), always)
		err = stacktrace.With(err, always)

		assert.Equal(t, `this is ERROR`, err.Error())

		assert.Contains(t,
			stacktrace.Format(err),
			"this is ERROR\n  [0]: ERROR\n\t",
		)

		assert.Equal(t, `this is ERROR`, fmt.Sprintf("%v", err))
		assert.Equal(t, `this is ERROR`, fmt.Sprintf("%+v", err))

		assert.Contains(t,
			fmt.Sprintf("%#v", err),
			`&stacktrace.Error{Cause:(*stacktrace.Error)(`,
		)

		assert.Len(t, stacktrace.Dump(err).Traces, 3)

		t.Log(stacktrace.Format(err))
	})

	t.Run("join", func(t *testing.T) {
		err0 := stacktrace.New("err0", stacktrace.Always)
		err1 := stacktrace.With(errors.New("err1"), stacktrace.Always)
		err1 = stacktrace.With(fmt.Errorf("%w(a)", err1), stacktrace.Always)
		err2 := stacktrace.New("err2", stacktrace.Always)
		err := errors.Join(err0, err1, err2)

		dump := stacktrace.Dump(err)
		assert.Len(t, dump.Traces, 4)

		assert.Contains(t, dump.Traces[0].StackEntries[0], "format_test.go:65")
		assert.Contains(t, dump.Traces[1].StackEntries[0], "format_test.go:66")
		assert.Contains(t, dump.Traces[2].StackEntries[0], "format_test.go:67")
		assert.Contains(t, dump.Traces[3].StackEntries[0], "format_test.go:68")

		t.Log(stacktrace.Format(err))
	})

	t.Run("consice format", func(t *testing.T) {
		err0 := stacktrace.New("err0")
		assert.NotContains(t, stacktrace.Format(err0), "[0]:")

		err1 := fmt.Errorf("err1: %w", err0)
		assert.Contains(t, stacktrace.Format(err1), "[0]:")
		assert.NotContains(t, stacktrace.Format(err1), "[1]:")

		err2 := stacktrace.With(err1, stacktrace.Always)
		assert.Contains(t, stacktrace.Format(err2), "[0]:")
		assert.Contains(t, stacktrace.Format(err2), "[1]:")
	})
}

func hello(options ...stacktrace.Option) error {
	return stacktrace.New("ERROR", options...)
}

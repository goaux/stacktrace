package stacktrace_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/goaux/stacktrace"
	"github.com/stretchr/testify/assert"
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

package stacktrace

import (
	"errors"
)

// New returns an error created with errors.New along with stack trace information in a concise way.
//
// Example usage:
//
//	err := stacktrace.New("something went wrong")
//
// This is equivalent to:
//
//	err := stacktrace.Trace(errors.New("something went wrong"))
//
// stacktrace.New is slightly more concise and efficient.
func New(text string) error {
	return newErrorSkip(errors.New(text), 1)
}

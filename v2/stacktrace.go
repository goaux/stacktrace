// Package stacktrace provides utilities for capturing and inspecting call
// stacks associated with errors.
//
// This package allows you to create error objects that include information
// about where in the code the error occurred, which can be extremely useful
// for debugging. It includes methods to unwrap underlying errors, retrieve
// call stack frames, and generate string representations of errors with
// detailed stack trace information.
package stacktrace

import (
	"errors"
	"fmt"
	"runtime"
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

// Errorf returns an error formatted according to the format specifier, similar to fmt.Errorf,
// along with stack trace information in a concise way.
//
// Example usage:
//
//	err := stacktrace.Errorf("failed to do %s", action)
//
// This is equivalent to:
//
//	err := stacktrace.Trace(fmt.Errorf("failed to do %s", action))
//
// stacktrace.Errorf is slightly more concise and efficient.
func Errorf(format string, a ...any) error {
	return withSkip(fmt.Errorf(format, a...), 1)
}

// Format returns a formatted string representation of the [DebugInfo] from err.
//
// If err is nil, it returns an empty string.
// If err's chain doesn't contain any stack trace information, it returns err.Error().
//
// This is equivalent to:
//
//	stacktrace.GetDebugInfo(err).Format()
//
// See [DebugInfo.Format].
func Format(err error) string {
	return GetDebugInfo(err).Format()
}

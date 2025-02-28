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

// Callers returns the program counters (PCs) of function invocations on the
// calling goroutine's stack, skipping the specified number of stack frames.
//
// The skip parameter determines how many initial stack frames to omit from the
// result. A skip value of 0 starts from the caller of Callers itself.
//
// The returned slice contains the collected program counters, which can be
// further processed using runtime.CallersFrames to obtain function details.
func Callers(skip int) []uintptr {
	skip += 2
	const size = 8
	var pc []uintptr
	for {
		x := make([]uintptr, size)
		n := runtime.Callers(skip, x)
		if n == 0 {
			break
		}
		pc = append(pc, x[:n]...)
		skip += n
	}
	return pc
}

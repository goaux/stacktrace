package stacktrace

import (
	"fmt"
	"runtime"
)

// Error wraps an error and adds a call stack indicating where it occurred.
//
// It is recommended to use the following functions to create an instance of Error:
//
//   - [Trace]: Adds a call stack indicating where [Trace] was executed to the given error and returns it.
//   - [Trace2], [Trace3], [Trace4]: Variants of [Trace] that accept multiple
//     arbitrary arguments along with an error and return them.
//     The trailing number represents the number of additional arguments other than the error.
//   - [Errorf]: A function that wraps `fmt.Errorf` with `Trace`.
//   - [New]: A function that wraps `errors.New` with `Trace`.
//
// If `nil` is passed to [Trace] and its variants, they return `nil` as an error.
//
// [Trace], [Trace2], [Trace3], [Trace4], and [Errorf] do not add a new call
// stack if the given error already has one, meaning that if an `Error`
// instance is already present in the error chain, the original error is
// returned as-is.
//
// [New] always returns an error with a newly added call stack.
type Error struct {
	// Err holds the original error that occurred.
	Err error

	// Callers contains the program counters (PCs) of the function call stack
	// at the time the error was captured. These can be used to retrieve stack traces.
	Callers []uintptr
}

var _ StackTracer = Error{}

// NewError returns a new Error instance with the given error and caller stack
// information.
func NewError(err error, callers []uintptr) *Error {
	return &Error{Err: err, Callers: callers}
}

func newErrorSkip(err error, skip int) error {
	return NewError(err, Callers(skip+1))
}

// Error returns a string representation of the custom error, including
// the original error message. If the call stack information
// is available, it appends the information about the first frame of
// the stack trace to the error message.
//
// The format of the returned string is:
// "<original error message> (<file name>:<line> <function name>)"
func (err Error) Error() string {
	var s string
	if err := err.Err; err != nil {
		s = err.Error()
	}
	if len(err.Callers) == 0 {
		return s
	}
	frame, _ := runtime.CallersFrames(err.Callers[:1]).Next()
	return fmt.Sprintf("%s (%s)", s, frameShortString(&frame))
}

// Unwrap returns the wrapped error, allowing for further inspection using
// errors.Unwrap.
func (err Error) Unwrap() error {
	return err.Err
}

// StackTrace returns a slice of program counters representing the call stack.
func (err Error) StackTrace() []uintptr {
	return err.Callers
}

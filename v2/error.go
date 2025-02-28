package stacktrace

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Error is a custom error type that captures the original error along with
// call stack information where the error was created.
//
// The following functions can be used to create and manipulate errors:
//
//   - [New]: Creates a new error instance without caller frames.
//   - [Errorf]: Creates an error using a format string, similar to `fmt.Errorf`.
//   - [Trace]: Wraps an existing error with call stack information.
//   - [Trace1], [Trace2], [Trace3]: Overloads of Trace that return the original values along with the traced error.
type Error struct {
	// Err is the underlying error value.
	Err error

	// Callers contains a list of stack frames at the time the error was created.
	Callers []uintptr
}

// NewError returns a new Error instance with the given error and caller stack
// information.
func NewError(err error, callers []uintptr) *Error {
	return &Error{Err: err, Callers: callers}
}

func extract(err error) *Error {
	var target *Error
	errors.As(err, &target)
	return target
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

func walkCallersFrames(pc []uintptr, fn func(*runtime.Frame)) {
	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		fn(&frame)
		if !more || frame.Function == "main.main" {
			break
		}
	}
}

func frameShortString(frame *runtime.Frame) string {
	return fmt.Sprintf(
		"%s:%d %s",
		filepath.Base(frame.File),
		frame.Line,
		frameFunction(frame.Function),
	)
}

func frameString(frame *runtime.Frame) string {
	return fmt.Sprintf(
		"%s:%d %s",
		frame.File,
		frame.Line,
		frameFunction(frame.Function),
	)
}

func frameFunction(s string) string {
	if i := strings.LastIndexByte(s, '/'); i != -1 {
		if j := strings.IndexByte(s[i+1:], '.'); i != -1 {
			s = s[i+j+2:]
		}
	}
	return s
}

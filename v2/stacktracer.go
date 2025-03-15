package stacktrace

import "errors"

// StackTracer represents an error that provides stack trace information.
type StackTracer interface {
	// StackTracer extends the error interface.
	error

	// StackTrace returns a slice of program counters representing the call stack.
	StackTrace() []uintptr
}

// HasStackTracer returns true if there is at least one StackTracer in the
// error chain, false otherwise.
func HasStackTracer(err error) bool {
	var other StackTracer
	return errors.As(err, &other)
}

// ListStackTracers returns all the StackTracers in the error chain.
func ListStackTracers(err error) []StackTracer {
	var list []StackTracer
	walkErrorChain(err, func(err error) {
		if v, ok := err.(StackTracer); ok {
			list = append(list, v)
		}
	})
	return list
}

func walkErrorChain(err error, callback func(err error)) {
	if err == nil {
		return
	}
	callback(err)
	switch err := err.(type) {
	case interface{ Unwrap() error }:
		walkErrorChain(err.Unwrap(), callback)
	case interface{ Unwrap() []error }:
		for _, err := range err.Unwrap() {
			walkErrorChain(err, callback)
		}
	}
}

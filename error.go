// Package stacktrace provides enhanced error handling capabilities with stack trace information.
//
// It offers a way to wrap errors with stack traces, extract stack information from error chains,
// and format error messages with detailed stack information.
package stacktrace

import (
	"errors"
	"fmt"
	"runtime"
)

// Error represents an error with an associated cause and stack frames.
type Error struct {
	Cause  error
	Frames []uintptr
}

// Error returns the error message of the underlying cause.
//
// It implements the error interface.
func (err Error) Error() string {
	return err.Cause.Error()
}

// Unwrap returns the underlying cause of the error.
//
// It allows Error to work with errors.Unwrap.
func (err Error) Unwrap() error {
	return err.Cause
}

// With wraps the given cause error with stack trace information.
//
// If cause is nil, it returns nil.
// Options can be provided to customize the behavior.
func With(cause error, options ...Option) error {
	if cause == nil {
		return nil
	}
	c := &config{
		Skip:  3,
		Limit: DefaultLimit,
	}
	c.apply(options...)
	return with(cause, c)
}

// New creates a new error with the given text and wraps it with stack trace information.
//
// It's meant to be used as a replacement for errors.New.
// Options can be provided to customize the behavior.
func New(text string, options ...Option) error {
	c := &config{
		Skip:  3,
		Limit: DefaultLimit,
	}
	c.apply(options...)
	return with(errors.New(text), c)
}

func with(cause error, c *config) error {
	if !c.Always {
		var v *Error
		if errors.As(cause, &v) {
			// If a *Error is already included in the cause,
			// do not wrap with another new *Error multiple times.
			return cause
		}
	}
	frames := make([]uintptr, c.Limit)
	n := runtime.Callers(c.Skip, frames)
	frames = frames[:n]
	return &Error{Cause: cause, Frames: frames}
}

// Errorf creates a new error using fmt.Errorf and wraps it with stack trace information.
//
// It's meant to be used as a replacement for fmt.Errorf.
// To specify options, use With(fmt.Errorf(...), options...) instead.
//
// see
//
//	stacktrace.Always.Errorf(...)
func Errorf(format string, a ...any) error {
	c := &config{
		Skip:  3,
		Limit: DefaultLimit,
	}
	return with(fmt.Errorf(format, a...), c)
}

// Extract returns a slice of *Error from the error chain of err.
//
// It traverses the error chain and collects all *Error instances.
// The returned slice is ordered such that the first added *Error
// (i.e., the one closest to the root cause) is at index 0,
// and the most recently added *Error (i.e., the one furthest from
// the root cause) is at the last index.
// This means that the stack trace added first in the error chain
// will be at the beginning of the returned slice.
// If err is nil or contains no *Error, an empty slice is returned.
func Extract(err error) []*Error {
	list := extract(nil, err)
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func extract(list []*Error, err error) []*Error {
	if v, ok := err.(*Error); ok {
		list = append(list, v)
	}
	switch v := err.(type) {
	case interface{ Unwrap() []error }:
		list = extractErrors(list, v.Unwrap())
	case interface{ Unwrap() error }:
		list = extract(list, v.Unwrap())
	}
	return list
}

func extractErrors(list []*Error, errs []error) []*Error {
	for i := len(errs) - 1; i >= 0; i-- {
		list = extract(list, errs[i])
	}
	return list
}

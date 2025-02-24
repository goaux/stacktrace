package stacktrace

// Trace returns the error along with stack frame information.
// It returns nil if the input error is nil.
func Trace(err error) error {
	return traceSkip(err, 1)
}

// Trace1 takes a value of any type and an error, returning the original value
// and nil if the error is not nil, or the error along with stack frame
// information.
func Trace1[T0 any](v0 T0, err error) (T0, error) {
	return v0, traceSkip(err, 1)
}

// Trace2 takes two values of any type and an error, returning the original
// values and nil if the error is not nil, or the error along with stack frame
// information.
func Trace2[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1, error) {
	return v0, v1, traceSkip(err, 1)
}

// Trace3 takes three values of any type and an error, returning the original
// values and nil if the error is not nil, or the error along with stack frame
// information.
func Trace3[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) (T0, T1, T2, error) {
	return v0, v1, v2, traceSkip(err, 1)
}

func traceSkip(err error, skip int) error {
	if err == nil {
		return nil
	}
	return withSkip(err, skip+1)
}

func withSkip(err error, skip int) error {
	if extract(err) != nil {
		return err
	}
	return newErrorSkip(err, skip+1)
}

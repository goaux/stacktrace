package stacktrace

// Trace returns the error along with stack frame information.
// It returns nil if the input error is nil.
func Trace(err error) error {
	return traceSkip(err, 1)
}

// Trace2 returns the given values unchanged if err is nil.
// If err is not nil, it wraps err with stack trace information before returning.
//
// The function name "Trace2" indicates that it takes two arguments
// (a value of any type and an error) and returns two values.
func Trace2[T0 any](v0 T0, err error) (T0, error) {
	return v0, traceSkip(err, 1)
}

// Trace3 returns the given values unchanged if err is nil.
// If err is not nil, it wraps err with stack trace information before returning.
//
// The function name "Trace3" indicates that it takes three arguments
// (two values of any type and an error) and returns three values.
func Trace3[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1, error) {
	return v0, v1, traceSkip(err, 1)
}

// Trace4 returns the given values unchanged if err is nil.
// If err is not nil, it wraps err with stack trace information before returning.
//
// The function name "Trace4" indicates that it takes four arguments
// (three values of any type and an error) and returns four values.
func Trace4[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) (T0, T1, T2, error) {
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

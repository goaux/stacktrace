package stacktrace

import "fmt"

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

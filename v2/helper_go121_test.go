//go:build go1.21

package stacktrace_test

import "slices"

// equal returns the result of calling `slices.Equal(s1, s2)`.
func equal(s1, s2 []uintptr) bool {
	return slices.Equal(s1, s2)
}

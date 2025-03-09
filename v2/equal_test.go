//go:build !go1.21

package stacktrace_test

// equal returns true if `s1` and `s2` are identical.
//
// In Go 1.20, [slices.Equal] was not available, so this custom implementation
// was created. Once support for Go versions 1.21 or later is ensured, it is
// recommended to use [slices.Equal` and remove this function.
func equal(s1, s2 []uintptr) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

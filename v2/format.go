package stacktrace

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

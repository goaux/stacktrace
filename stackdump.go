package stacktrace

// StackDump represents a structured version of an error chain containing *Error.
type StackDump struct {
	// Error stores the result of err.Error().
	Error string `json:"error,omitempty"`

	// Traces stores the result of Extract(err) converted to StackTrace.
	Traces []StackTrace `json:"traces,omitempty"`
}

// StackTrace represents a structured version of a single *Error in an error chain.
type StackTrace struct {
	// Detail stores the result of Error.Cause.Error().
	Detail string `json:"detail,omitempty"`

	// StackEntries stores the result of Frames(Error.Frames).
	StackEntries []string `json:"stack_entries,omitempty"`
}

// Dump converts the result of Extract(err) to a StackDump.
// It provides a structured representation equivalent to what Format returns as a string.
// Dump returns the zero value if err is nil.
func Dump(err error) StackDump {
	if err == nil {
		return StackDump{}
	}
	dump := StackDump{Error: err.Error()}
	list := Extract(err)
	if n := len(list); n > 0 {
		dump.Traces = make([]StackTrace, n)
		for i, e := range list {
			dump.Traces[i].Detail = e.Cause.Error()
			dump.Traces[i].StackEntries = Frames(e.Frames)
		}
	}
	return dump
}

package stacktrace

import (
	"runtime"
)

// DebugInfo represents debug information about an error.
//
// This struct is compatible with [google.golang.org/genproto/googleapis/rpc/errdetails.DebugInfo].
type DebugInfo struct {
	// StackEntries contains a list of stack trace entries related to the error.
	StackEntries []string `json:"stack_entries,omitempty"`

	// Detail provides a detailed error message.
	Detail string `json:"detail,omitempty"`
}

// GetDebugInfo extracts debug information from an error.
// It collects stack trace frames and formats them as strings, then returns
// this information along with the detailed error message in a [DebugInfo] struct.
//
// It returns a zero value if err is nil.
func GetDebugInfo(err error) DebugInfo {
	if err == nil {
		return DebugInfo{}
	}
	var frames []string
	if err := extract(err); err != nil {
		frames = make([]string, 0, len(err.Callers))
		walkCallersFrames(err.Callers, func(frame *runtime.Frame) {
			frames = append(frames, frameString(frame))
		})
	}
	detail := err.Error()
	return DebugInfo{Detail: detail, StackEntries: frames}
}

package stacktrace

import (
	"runtime"
	"strings"
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

// Format returns a formatted string representation of the DebugInfo.
//
// The output consists of the Detail message followed by the stack trace entries,
// each separated by a newline and a tab ("\n\t").
func (info DebugInfo) Format() string {
	n := len(info.StackEntries)
	if n == 0 {
		return info.Detail
	}
	return strings.Join(
		append(append(make([]string, 0, 1+n), info.Detail), info.StackEntries...),
		"\n\t",
	)
}

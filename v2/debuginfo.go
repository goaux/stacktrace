package stacktrace

import (
	"runtime"
	"strings"
)

// DebugInfo represents debug information about an error.
//
// This struct is compatible with [google.golang.org/genproto/googleapis/rpc/errdetails.DebugInfo].
type DebugInfo struct {
	// Detail provides a detailed error message.
	Detail string `json:"detail,omitempty"`

	// StackEntries contains a list of stack trace entries related to the error.
	StackEntries []string `json:"stack_entries,omitempty"`
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
	return DebugInfo{
		Detail:       err.Error(),
		StackEntries: stackEntries(err),
	}
}

func stackEntries(err error) []string {
	list := ListStackTracers(err)
	if len(list) == 0 {
		return nil
	}
	detail := err.Error()
	var entries []string
	for i, v := range list {
		if i != 0 || detail != v.Error() {
			entries = append(entries, "## "+v.Error())
		}
		callers := v.StackTrace()
		walkCallersFrames(callers, func(frame *runtime.Frame) {
			entries = append(entries, frameString(frame))
		})
	}
	return entries
}

func walkCallersFrames(pc []uintptr, fn func(*runtime.Frame)) {
	if len(pc) == 0 {
		return
	}
	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		fn(&frame)
		if !more || frame.Function == "main.main" {
			break
		}
	}
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

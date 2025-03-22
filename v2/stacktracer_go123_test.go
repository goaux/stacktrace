//go:build go1.23

package stacktrace_test

import "github.com/goaux/stacktrace/v2"

func ExampleListStackTracers() {
	var err error
	for _, v := range stacktrace.ListStackTracers(err) {
		for frame := range stacktrace.CallersFrames(v.StackTrace()) {
			_, _ = frame.File, frame.Line
		}
	}
}

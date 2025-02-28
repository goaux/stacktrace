//go:build go1.23

package stacktrace

import (
	"iter"
	"runtime"
)

// CallersFrames returns an iterator sequence of *[runtime.Frame] based on the given
// slice of program counters (PCs). It yields stack frames one by one, stopping
// if the iterator function returns false, if there are no more frames, or if
// the function name is "main.main".
//
// The function utilizes runtime.CallersFrames to traverse the stack trace
// and provides each frame through the yield function.
func CallersFrames(callers []uintptr) iter.Seq[*runtime.Frame] {
	return func(yield func(*runtime.Frame) bool) {
		frames := runtime.CallersFrames(callers)
		for {
			frame, more := frames.Next()
			if !yield(&frame) || !more {
				break
			}
		}
	}
}

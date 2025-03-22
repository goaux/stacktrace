package stacktrace

import "runtime"

// Callers returns the program counters (PCs) of function invocations on the
// calling goroutine's stack, skipping the specified number of stack frames.
//
// The skip parameter determines how many initial stack frames to omit from the
// result. A skip value of 0 starts from the caller of Callers itself.
//
// The returned slice contains the collected program counters, which can be
// further processed using runtime.CallersFrames to obtain function details.
func Callers(skip int) []uintptr {
	skip += 2
	const size = 16
	var pc []uintptr
	for {
		x := make([]uintptr, size)
		n := runtime.Callers(skip, x)
		if n != len(x) {
			pc = append(pc, x[:n]...)
			break
		}
		pc = append(pc, x...)
		skip += n
	}
	return pc
}

package stacktrace

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func frameString(frame *runtime.Frame) string {
	return fmt.Sprintf(
		"%s:%d %s",
		frame.File,
		frame.Line,
		frameFunction(frame.Function),
	)
}

func frameShortString(frame *runtime.Frame) string {
	return fmt.Sprintf(
		"%s:%d %s",
		filepath.Base(frame.File),
		frame.Line,
		frameFunction(frame.Function),
	)
}

func frameFunction(s string) string {
	if i := strings.LastIndexByte(s, '/'); i != -1 {
		if j := strings.IndexByte(s[i+1:], '.'); i != -1 {
			s = s[i+j+2:]
		}
	}
	return s
}

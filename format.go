package stacktrace

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Format returns a string representation of the error, including stack trace information if available.
//
// If err is nil, it returns an empty string.
// If err's chain doesn't contain any *Error, it returns err.Error().
func Format(err error) string {
	if err == nil {
		return ""
	}
	list := Extract(err)
	if len(list) == 0 {
		return err.Error()
	}
	s := &strings.Builder{}
	s.WriteString(err.Error())
	for i, e := range list {
		if 0 < i || e.Error() != err.Error() {
			s.WriteString("\n  [")
			s.WriteString(strconv.Itoa(i))
			s.WriteString("]: ")
			s.WriteString(e.Error())
		}
		for _, frame := range Frames(e.Frames) {
			s.WriteString("\n\t")
			s.WriteString(frame)
		}
	}
	return s.String()
}

// Frames converts a slice of stack frame pointers to a slice of formatted strings.
//
// Each string in the returned slice has the format "<file>:<line> <function>".
func Frames(frames []uintptr) []string {
	lines := make([]string, 0, len(frames))
	iter := runtime.CallersFrames(frames)
	for {
		fr, more := iter.Next()
		fn := funcname(fr.Function)
		line := fmt.Sprintf("%s:%d %s", fr.File, fr.Line, fn)
		lines = append(lines, line)
		if !more || fn == "main" {
			break
		}
	}
	return lines
}

func funcname(name string) string {
	if i := strings.LastIndexByte(name, '/'); i >= 0 {
		if j := strings.IndexByte(name[i+1:], '.'); j >= 0 {
			return name[i+j+2:]
		}
		return name[i+1:]
	}
	if i := strings.IndexByte(name, '.'); i >= 0 {
		return name[i+1:]
	}
	return name
}

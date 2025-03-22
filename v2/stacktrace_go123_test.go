//go:build go1.23

package stacktrace_test

import (
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestCallersFrames(t *testing.T) {
	f1 := func() []uintptr { return stacktrace.Callers(0) }
	f2 := func() []uintptr { return f1() }
	f3 := func() []uintptr { return f2() }
	callers := f3()

	list := make([]string, 0, len(callers))
	for frame := range stacktrace.CallersFrames(callers) {
		list = append(list, funcnameOf(t, frame.PC))
	}

	if got, want := len(list), len(callers); got != want {
		t.Errorf("len(list)=%d must be len(callers)=%d", got, want)
	}

	want := stacktrace.GetDebugInfo(stacktrace.NewError(nil, callers)).StackEntries
	for i, s := range want {
		want[i] = strings.Fields(s)[1]
		if j := strings.Index(list[i], want[i]); j != -1 {
			list[i] = list[i][j:]
		}
	}
	if !slices.Equal(list, want) {
		t.Errorf("list must be equal to callers; list=%v callers=%v", list, callers)
	}

	t.Run("nil", func(t *testing.T) {
		n := 0
		for range stacktrace.CallersFrames(nil) {
			n++
		}
		if n != 0 {
			t.Errorf("n must be 0, but %d", n)
		}
	})
}

func funcnameOf(t *testing.T, pc uintptr) string {
	t.Helper()
	f := runtime.FuncForPC(pc)
	if f == nil {
		t.Fatal("runtime.FuncForPC(pc) must not return nil")
	}
	return f.Name()
}

func ExampleListStackTracers() {
	var err error
	for _, v := range stacktrace.ListStackTracers(err) {
		for frame := range stacktrace.CallersFrames(v.StackTrace()) {
			_, _ = frame.File, frame.Line
		}
	}
}

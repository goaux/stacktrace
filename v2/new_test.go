package stacktrace_test

import (
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestNew(t *testing.T) {
	err := stacktrace.New("42")
	info := stacktrace.GetDebugInfo(err)
	want := "42 (new_test.go:10 TestNew)"
	if got := info.Detail; got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
	if n := len(info.StackEntries); n != 3 {
		t.Errorf("len(info.StackEntries) = %d, must be 3", n)
	}
}

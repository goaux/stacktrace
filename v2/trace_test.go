package stacktrace_test

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestTrace(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		ok := func() error { return nil }
		err := stacktrace.Trace(ok())
		if err != nil {
			t.Error("err must be nil")
		}
	})
	t.Run("StackEntries[0]", func(t *testing.T) {
		err := stacktrace.Trace(os.Chdir("/no/such/dir"))
		_, file, line, _ := runtime.Caller(0)
		debugInfo := stacktrace.GetDebugInfo(err)
		want := fmt.Sprintf("%s:%d TestTrace.func2", file, line-1)
		if got := debugInfo.StackEntries[0]; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
	})
}

func TestTrace1(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		ok := func() (int, error) { return 42, nil }
		i, err := stacktrace.Trace1(ok())
		if err != nil {
			t.Error("err must be nil")
		}
		if i != 42 {
			t.Error("i must be 42")
		}
	})
	t.Run("StackEntries[0]", func(t *testing.T) {
		_, err := stacktrace.Trace1(os.ReadDir("/no/such/dir"))
		_, file, line, _ := runtime.Caller(0)
		debugInfo := stacktrace.GetDebugInfo(err)
		want := fmt.Sprintf("%s:%d TestTrace1.func2", file, line-1)
		if got := debugInfo.StackEntries[0]; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
	})
}

func TestTrace2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		ok := func() (int, string, error) { return 42, "42", nil }
		i, s, err := stacktrace.Trace2(ok())
		if err != nil {
			t.Error("err must be nil")
		}
		if i != 42 {
			t.Error("i must be 42")
		}
		if s != "42" {
			t.Error("s must be 42")
		}
	})
	t.Run("StackEntries[0]", func(t *testing.T) {
		fn := func() (s string, i int, err error) {
			err = errors.New("TESTING")
			return
		}
		_, _, err := stacktrace.Trace2(fn())
		_, file, line, _ := runtime.Caller(0)
		debugInfo := stacktrace.GetDebugInfo(err)
		want := fmt.Sprintf("%s:%d TestTrace2.func2", file, line-1)
		if got := debugInfo.StackEntries[0]; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
	})
}

func TestTrace3(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		ok := func() (int, string, float64, error) { return 42, "42", 42.42, nil }
		i, s, f, err := stacktrace.Trace3(ok())
		if err != nil {
			t.Error("err must be nil")
		}
		if i != 42 {
			t.Error("i must be 42")
		}
		if s != "42" {
			t.Error("s must be 42")
		}
		if f != 42.42 {
			t.Error("f must be 42.42")
		}
	})
	t.Run("StackEntries[0]", func(t *testing.T) {
		fn := func() (s string, i int, b bool, err error) {
			err = errors.New("TESTING")
			return
		}
		_, _, _, err := stacktrace.Trace3(fn())
		_, file, line, _ := runtime.Caller(0)
		debugInfo := stacktrace.GetDebugInfo(err)
		want := fmt.Sprintf("%s:%d TestTrace3.func2", file, line-1)
		if got := debugInfo.StackEntries[0]; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
	})
}

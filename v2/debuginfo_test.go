package stacktrace_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func TestGetDebugInfo(t *testing.T) {
	t.Run("", func(t *testing.T) {
		err := stacktrace.Trace(os.Chdir("/no/such/dir"))
		_, file, line, _ := runtime.Caller(0)
		info := stacktrace.GetDebugInfo(err)
		want := fmt.Sprintf(
			`chdir /no/such/dir: no such file or directory (%s:%d TestGetDebugInfo.func1)`,
			filepath.Base(file), line-1,
		)
		if got := info.Detail; got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
		if len(info.StackEntries) == 0 {
			t.Errorf("len(info.StackEntries) must be greater than 0")
		}
	})

	t.Run("", func(t *testing.T) {
		info := stacktrace.GetDebugInfo(os.ErrDeadlineExceeded)
		want := os.ErrDeadlineExceeded.Error()
		if got := info.Detail; got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
		if len(info.StackEntries) != 0 {
			t.Errorf("len(info.StackEntries) must be 0")
		}
	})
}

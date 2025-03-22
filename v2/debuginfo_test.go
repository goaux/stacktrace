package stacktrace_test

import (
	"errors"
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

	t.Run("with no stack trace", func(t *testing.T) {
		err := stacktrace.NewError(errors.New("testing"), nil)
		got := stacktrace.GetDebugInfo(err).Format()
		want := err.Error()
		if got != want {
			t.Errorf("got=%q want=%q", got, want)
		}
	})
}

func TestDebugInfo_Format(t *testing.T) {
	tests := []struct {
		name string
		info stacktrace.DebugInfo
		want string
	}{
		{
			name: "full",
			info: stacktrace.DebugInfo{StackEntries: []string{"1", "2"}, Detail: "detail"},
			want: "detail\n\t1\n\t2",
		},
		{
			name: "stackentries",
			info: stacktrace.DebugInfo{StackEntries: []string{"1", "2"}},
			want: "\n\t1\n\t2",
		},
		{
			name: "detail",
			info: stacktrace.DebugInfo{Detail: "detail"},
			want: "detail",
		},
		{
			name: "zero",
			info: stacktrace.DebugInfo{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.info.Format()
			if got != tt.want {
				t.Errorf("got=%s want=%s", got, tt.want)
			}
		})
	}
}

func ExampleDebugInfo_Format() {
	err := stacktrace.Trace(os.Chdir("/no/such/dir"))
	fmt.Println(stacktrace.Format(err))
}

func ExampleDebugInfo_Format_nil() {
	err := stacktrace.Trace(nil)
	fmt.Println(stacktrace.Format(err))
	// Output:
}

func ExampleDebugInfo_multiple() {
	ch := make(chan error)
	go func() { ch <- stacktrace.New("hello") }()
	go func() { ch <- stacktrace.New("world") }()
	err1 := <-ch // error from the thread
	err2 := <-ch // error from the other thread
	if err1 != nil || err2 != nil {
		err := errors.Join(
			stacktrace.New("Two errors"), // error of this thread
			err1,
			err2,
		)
		info := stacktrace.GetDebugInfo(err)
		fmt.Println(info.Format()) // this prints 3 stack frames.
	}
}

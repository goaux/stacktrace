package stacktrace_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/goaux/stacktrace/v2"
)

func Example() {
	if err := stacktrace.Trace(os.Chdir(".")); err != nil {
		fmt.Println("err.Error():", err.Error())
		info := stacktrace.GetDebugInfo(err)
		fmt.Println("info.Detail:", info.Detail)
		fmt.Printf("%d stack entries included.", len(info.StackEntries))
		panic(stacktrace.GetDebugInfo(err))
	}
	if file, err := stacktrace.Trace1(os.Open("./no/such/file")); err != nil {
		fmt.Println("err.Error():", err.Error())
		info := stacktrace.GetDebugInfo(err)
		fmt.Println("info.Detail:", info.Detail)
		fmt.Printf("%d stack entries included.", len(info.StackEntries))
	} else {
		file.Close()
	}
	// Output:
	// err.Error(): open ./no/such/file: no such file or directory (stacktrace_test.go:19 Example)
	// info.Detail: open ./no/such/file: no such file or directory (stacktrace_test.go:19 Example)
	// 5 stack entries included.
}

func TestNew(t *testing.T) {
	err := stacktrace.New("42")
	info := stacktrace.GetDebugInfo(err)
	want := "42 (stacktrace_test.go:34 TestNew)"
	if got := info.Detail; got != want {
		t.Errorf("want=%q got=%q", want, got)
	}
	if n := len(info.StackEntries); n != 3 {
		t.Errorf("len(info.StackEntries) = %d, must be 3", n)
	}
}

func TestErrorf(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		err := stacktrace.Errorf("42 %w", os.ErrInvalid)
		info := stacktrace.GetDebugInfo(err)
		want := "42 invalid argument (stacktrace_test.go:47 TestErrorf.func1)"
		if got := info.Detail; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
		if n := len(info.StackEntries); n != 3 {
			t.Errorf("len(info.StackEntries) = %d, must be 3", n)
		}
	})
	t.Run("single", func(t *testing.T) {
		err := stacktrace.Errorf("42 %w", os.ErrInvalid)
		err = stacktrace.Errorf("41 %w", err)
		err = stacktrace.Errorf("40 %w", err)
		info := stacktrace.GetDebugInfo(err)
		want := "40 41 42 invalid argument (stacktrace_test.go:58 TestErrorf.func2)"
		if got := info.Detail; got != want {
			t.Errorf("want=%q got=%q", want, got)
		}
		if n := len(info.StackEntries); n != 3 {
			t.Errorf("len(info.StackEntries) = %d, must be 3", n)
		}
	})
}

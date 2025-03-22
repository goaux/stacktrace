package stacktrace_test

import (
	"fmt"
	"os"

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
	if file, err := stacktrace.Trace2(os.Open("./no/such/file")); err != nil {
		fmt.Println("err.Error():", err.Error())
		info := stacktrace.GetDebugInfo(err)
		fmt.Println("info.Detail:", info.Detail)
		fmt.Printf("%d stack entries included.", len(info.StackEntries))
	} else {
		file.Close()
	}
	// Output:
	// err.Error(): open ./no/such/file: no such file or directory (stacktrace_test.go:18 Example)
	// info.Detail: open ./no/such/file: no such file or directory (stacktrace_test.go:18 Example)
	// 5 stack entries included.
}

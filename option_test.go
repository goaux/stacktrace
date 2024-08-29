package stacktrace_test

import (
	"fmt"
	"testing"

	"github.com/goaux/stacktrace"
)

func ExampleAlways() {
	err0 := stacktrace.New("this is ERROR (0)")
	err1 := stacktrace.Always.Errorf("err1: %w", err0)
	list := stacktrace.Extract(err1)

	fmt.Println("err1.Error() = " + err1.Error())
	fmt.Printf("len = %d\n", len(list))
	fmt.Println(list[0].Cause.Error())
	fmt.Println(list[1].Cause.Error())
	fmt.Println()

	err2 := stacktrace.Errorf("err2: %w", err0)
	list = stacktrace.Extract(err2)

	fmt.Println("err2.Error() = " + err2.Error())
	fmt.Printf("len = %d\n", len(list))
	fmt.Println(list[0].Cause.Error())

	// Output:
	// err1.Error() = err1: this is ERROR (0)
	// len = 2
	// this is ERROR (0)
	// err1: this is ERROR (0)
	//
	// err2.Error() = err2: this is ERROR (0)
	// len = 1
	// this is ERROR (0)
}

func TestOption(t *testing.T) {
	stacktrace.Errorf("this is ERROR")
	stacktrace.Always.Errorf("this is ERROR")
}

func TestForce(t *testing.T) {
	err0 := stacktrace.New("this is ERROR")

	err1 := stacktrace.Always.Errorf("error: %w", err0)
	v0 := stacktrace.Extract(err1)
	assert.Len(t, v0, 2)

	err2 := stacktrace.Errorf("error: %w", err0)
	v1 := stacktrace.Extract(err2)
	assert.Len(t, v1, 1)
}

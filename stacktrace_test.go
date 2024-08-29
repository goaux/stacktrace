package stacktrace_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/goaux/stacktrace"
)

func Example() {
	err := doSomethingRisky()
	if err != nil {
		fmt.Println("An error occurred:")
		fmt.Println("error: " + stacktrace.Format(err))

		fmt.Println("\nJSON representation of the error:")
		dumpJSON, _ := json.MarshalIndent(stacktrace.Dump(err), "", "  ")
		fmt.Println(string(dumpJSON))
	}
}

func doSomethingRisky() error {
	err := openNonExistentFile()
	if err != nil {
		return stacktrace.With(err, stacktrace.Skip(1))
	}
	return nil
}

func openNonExistentFile() error {
	_, err := os.Open("non_existent_file.txt")
	if err != nil {
		return stacktrace.Errorf("failed to open file: %w", err)
	}
	return nil
}

type asserter struct{}

var assert asserter

func (asserter) Contains(t *testing.T, s, contains string) bool {
	t.Helper()
	r := strings.Contains(s, contains)
	if !r {
		t.Errorf("Error: %q does not contain %q", s, contains)
	}
	return r
}

func (asserter) NotContains(t *testing.T, s, contains string) bool {
	t.Helper()
	r := strings.Contains(s, contains)
	if r {
		t.Errorf("Error: %q should not contain %q", s, contains)
	}
	return r
}

func (asserter) Equal(t *testing.T, expected, actual any) bool {
	t.Helper()
	r := reflect.DeepEqual(expected, actual)
	if !r {
		t.Errorf("Error: Not equal: expected: %#v actual: %#v", expected, actual)
	}
	return r
}

func (asserter) Len(t *testing.T, object any, length int) bool {
	t.Helper()
	rv := reflect.ValueOf(object)
	r := rv.Len() == length
	if !r {
		t.Errorf("Error: %#v should have %d item(s), but has %d", object, length, rv.Len())
	}
	return r
}

func (asserter) True(t *testing.T, value bool) bool {
	t.Helper()
	if !value {
		t.Error("Error: should be true")
	}
	return value
}

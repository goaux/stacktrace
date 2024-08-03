package stacktrace_test

import (
	"encoding/json"
	"fmt"
	"os"

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

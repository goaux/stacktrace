# stacktrace/v2

Package stacktrace provides utilities for capturing and inspecting call stacks associated with errors.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/stacktrace/v2.svg)](https://pkg.go.dev/github.com/goaux/stacktrace/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/stacktrace/v2)](https://goreportcard.com/report/github.com/goaux/stacktrace/v2)

## Features

- Wraps errors with a single call stack
  - Typically, a single error chain contains at most one call stack
- Extracts the call stack from an error

## Usage

### Trace

The most basic way to create an error with a stack trace is to use the `Trace` function:

```go
err := stacktrace.Trace(os.Chdir(target))
```

There are overloads of `Trace` that return the original values along with the error.
These are `Trace2`, `Trace3`, and `Trace4`:

```go
file, err := stacktrace.Trace2(os.Open(file))
```

### New and Errorf

For convenience, you can use `New` and `Errorf` as drop-in replacements for `errors.New` and `fmt.Errorf`:

```go
err := stacktrace.New("some error")
err := stacktrace.Errorf("some error: %w", originalErr)
```

## Extracting Call Stack Information

### As a string

To get a formatted string representation of call stack information from an error:

```go
// Get as a string.
// This is equivalent to calling `err.Error()`
// if `err` does not contain call stack information.
// `s` is an empty string if `err` is nil.
s := stacktrace.Format(err)
```

For example, the `s` contains multiline string like below:

```
chdir /no/such/dir: no such file or directory (debuginfo_test.go:80 ExampleDebugInfo_Format)
        github.com/goaux/stacktrace/v2/debuginfo_test.go:80 ExampleDebugInfo_Format
        testing/run_example.go:63 testing.runExample
        testing/example.go:41 testing.runExamples
        testing/testing.go:2144 testing.(*M).Run
        _testmain.go:73 main.main
```

### As a DebugInfo

To extract call stack information from an error:

```go
// Get as a DebugInfo instance
info := stacktrace.GetDebugInfo(err)
```

The [DebugInfo](https://pkg.go.dev/github.com/goaux/stacktrace/v2#DebugInfo) type is compatible with [google.golang.org/genproto/googleapis/rpc/errdetails.DebugInfo](https://pkg.go.dev/google.golang.org/genproto/googleapis/rpc/errdetails#DebugInfo).

For example, the info contains information like below:

```
{
  "stack_entries": [
    "github.com/goaux/stacktrace/v2/debuginfo_test.go:81 ExampleDebugInfo_Format",
    "testing/run_example.go:63 testing.runExample",
    "testing/example.go:41 testing.runExamples",
    "testing/testing.go:2144 testing.(*M).Run",
    "_testmain.go:73 main.main"
  ],
  "detail": "chdir /no/such/dir: no such file or directory (debuginfo_test.go:81 ExampleDebugInfo_Format)"
}
```

### Other ways

Alternatively, you can use `errors.As` to extract the `Error` instance from an error chain.

## Performance Considerations

Adding stack traces to errors involves some overhead. In performance-critical
sections, consider using traditional error handling and adding stack traces at
higher levels of your application.

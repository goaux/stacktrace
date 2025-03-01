# stacktrace/v2

Package stacktrace provides utilities for capturing and inspecting stack traces associated with errors.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/stacktrace/v2.svg)](https://pkg.go.dev/github.com/goaux/stacktrace/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/stacktrace/v2)](https://goreportcard.com/report/github.com/goaux/stacktrace/v2)

## Features

- Wraps errors with a single stack trace
  - Typically, a single error chain contains at most one stack trace
- Extracts the stack trace from an error

## Usage

### Trace

The most basic way to create an error with a stack trace is to use the `Trace` function:

```go
err := stacktrace.Trace(os.Chdir(target))
```

There are `Trace2`, `Trace3`, and `Trace4`:
These are overloads of `Trace` that return the given values unchanged if err is nil.
If err is not nil, it wraps err with stack trace information before returning.

```go
file, err := stacktrace.Trace2(os.Open(file))
```

### New and Errorf

For convenience, you can use `New` and `Errorf` as drop-in replacements for `errors.New` and `fmt.Errorf`:

```go
err := stacktrace.New("some error")
err := stacktrace.Errorf("some error: %w", originalErr)
```

## Extracting Stack Trace Information

### As a string

To get a formatted string representation of stack trace information from an error:

```go
// Get as a string.
// This is equivalent to calling `err.Error()`
// if `err` does not contain stack trace information.
// `s` is an empty string if `err` is nil.
s := stacktrace.Format(err)
```

For example, the `s` contains multiline string like below:

```
chdir /no/such/dir: no such file or directory (run.go:10 main.run)
        example.com/hello/run.go:10 main.run
        example.com/hello/main.go:11 main.main
```

### As a DebugInfo

To extract stack trace information from an error:

```go
// Get as a DebugInfo instance
info := stacktrace.GetDebugInfo(err)
```

The [DebugInfo](https://pkg.go.dev/github.com/goaux/stacktrace/v2#DebugInfo) type is compatible with [google.golang.org/genproto/googleapis/rpc/errdetails.DebugInfo](https://pkg.go.dev/google.golang.org/genproto/googleapis/rpc/errdetails#DebugInfo).

For example, the info contains information like below:

```
{
  "detail": "chdir /no/such/dir: no such file or directory (run.go:10 main.run)",
  "stack_entries": [
    "example.com/hello/run.go:10 main.run",
    "example.com/hello/main.go:11 main.main"
  ]
}
```

### Other ways

Alternatively, you can use `errors.As` to extract the `Error` instance from an error chain.

## Performance Considerations

Adding stack traces to errors involves some overhead. In performance-critical
sections, consider using traditional error handling and adding stack traces at
higher levels of your application.

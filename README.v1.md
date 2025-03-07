# stacktrace
Package stacktrace provides enhanced error handling capabilities with call stack dumps for Go applications.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/stacktrace.svg)](https://pkg.go.dev/github.com/goaux/stacktrace)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/stacktrace)](https://goreportcard.com/report/github.com/goaux/stacktrace)

## Features

- Wrap errors with call stack information
- Extract and format stack traces from errors
- Support for error chains
- JSON-serializable error representations
- Customizable stack trace depth and skip frames

## Usage

### Creating Errors with Stack Traces

The most basic way to create an error with a stack trace is to use the `With` function:

    err := stacktrace.With(errors.New("some error"))

For convenience, you can use `New` and `Errorf` as drop-in replacements for `errors.New` and `fmt.Errorf`:

    err := stacktrace.New("some error")
    err := stacktrace.Errorf("some error: %w", originalErr)

### Extracting Stack Traces

To extract stack trace information from an error:

    // Get as a formatted string
    formattedError := stacktrace.Format(err)

    // Get as a structured StackDump
    dumpedError := stacktrace.Dump(err)

    // Get as a slice of *Error
    errorSlice := stacktrace.Extract(err)

### Working with Error Chains

By default, `With` only adds a stack trace to the first error in the chain:

    err1 := stacktrace.With(errors.New("first error"))
    err2 := stacktrace.With(err1) // err2 is the same as err1

To force adding multiple stack traces in an error chain, use the `Always` option:

    err1 := stacktrace.With(errors.New("first error"))
    err2 := stacktrace.With(err1, stacktrace.Always) // err2 has two stack traces

For `Errorf` with `Always`:

    err2 := stacktrace.Always.Errorf("second error: %w", err1)

### Customizing Stack Traces

You can customize the depth of the stack trace and skip frames:

    err := stacktrace.With(originalErr, stacktrace.Limit(10), stacktrace.Skip(1))

## Advanced Usage

### JSON Serialization

The `StackDump` struct returned by `Dump` is JSON-serializable:

    dump := stacktrace.Dump(err)
    jsonData, _ := json.Marshal(dump)

### Integration with Logging

You can easily integrate stacktrace with your logging system:

    if err != nil {
        log.Printf("Error occurred: %s", stacktrace.Format(err))
    }

### About Multiple Call Stack Dumps

Errors are returned from deep in the call stack towards the shallow direction.
Call stack dumps obtained from deep locations also include information from
shallower locations. Therefore, normally, having just one call stack dump in
the error chain contains the necessary information.

By default, `With` only adds a stack trace to the first error in the chain:

    err0 := errors.New("error message")
    err1 := stacktrace.With(err0) // err1 is an instance of *Error
    err2 := stacktrace.With(err1) // Since err1 already has a stack dump set, it returns err1

However, in cases where an error occurring in one goroutine is passed to
another goroutine to propagate the error, you might want to set multiple call
stack dumps in the error chain. For this purpose, use the `stacktrace.Always`
option:

    err0 := errors.New("error message")
    err1 := stacktrace.With(err0) // err1 is an instance of *Error
    err3 := stacktrace.With(err1, stacktrace.Always) // err3 has two call stack dumps

    // For Errorf with Always
    err4 := stacktrace.Always.Errorf("work error %w", err3) // err4 has three call stack dumps

## Best Practices

1. Use `stacktrace.New` and `stacktrace.Errorf` at the point where errors originate.
2. Use `stacktrace.With` when wrapping errors from external packages.
3. Use `stacktrace.Always` when you need to track error propagation across goroutines.
4. Always check if an error is nil before using stacktrace functions.

## Performance Considerations

Adding stack traces to errors involves some overhead. In performance-critical
sections, consider using traditional error handling and add stack traces at
higher levels of your application.

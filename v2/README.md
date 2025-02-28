# stacktrace/v2

Package stacktrace provides utilities for capturing and inspecting call stacks associated with errors.

## Features

- Wraps errors with a single call stack
- Extracts the call stack from an error

## Usage

The most basic way to create an error with a stack trace is to use the `Trace` function:

```go
err := stacktrace.Trace(os.Chdir(target))
```

There are overlods of `Trace` that return the original values along with the error.
They are `Trace1`, `Trace2` and `Trace3`.

```go
file, err := stacktrace.Trace1(os.Open(file))
```

For convenience, you can use `New` and `Errorf` as drop-in replacements for `errors.New` and `fmt.Errorf`:

```go
err := stacktrace.New("some error")
err := stacktrace.Errorf("some error: %w", originalErr)
```

## Extracting Call Stack Information

To get a formatted string representation of call stack information from an error:

```go
// Get as a string.
// This is equivalent to the result of calling `err.Error()`
// if err doesn't conatin call stack information.
// s is an empty string if err is nil.
s := stacktrace.Format(err)
```

To extract call stack information from an error:

```go
// Get as a DebugInfo
info := stacktrace.GetDebugInfo(err)
```

Or you can use `errors.As` to extract the Error instance from an error chain.

## Performance Considerations

Adding stack traces to errors involves some overhead. In performance-critical
sections, consider using traditional error handling and add stack traces at
higher levels of your application.

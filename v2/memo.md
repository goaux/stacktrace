---
title: memo
date: 2025-03-22T15:22:15.501+09:00
---

```
func Callers(skip int) []uintptr
func CallersFrames(callers []uintptr) iter.Seq[*runtime.Frame]
func Errorf(format string, a ...any) error
func Format(err error) string
func HasStackTracer(err error) bool
func New(text string) error
func Trace(err error) error
func Trace2[T0 any](v0 T0, err error) (T0, error)
func Trace3[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1, error)
func Trace4[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) (T0, T1, T2, error)
type DebugInfo struct{ ... }
    func GetDebugInfo(err error) DebugInfo
type Error struct{ ... }
    func NewError(err error, callers []uintptr) *Error
type StackTracer interface{ ... }
    func ListStackTracers(err error) []StackTracer
```

## BEFORE

debuginfo.go
error.go
frame.go
stacktrace.go
stacktracer.go
trace.go

## AFTER

callers.go
debuginfo.go DebugInfo
error.go Error
errorf.go
format.go
frame.go
helper_go121_test.go
helper_legacy_test.go
new.go
stacktracer.go HasStackTracer StackTracer
trace.go Trace Trace2 Trace3 Trace4

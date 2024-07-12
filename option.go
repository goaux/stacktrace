package stacktrace

import "fmt"

// DefaultLimit is the default depth for retrieving stack frames.
const DefaultLimit = 32

// Option is the type for options that modify the behavior of With.
type Option interface {
	apply(*config)
}

type config struct {
	Skip   int
	Limit  int
	Always bool
}

func (c *config) apply(options ...Option) {
	for _, o := range options {
		o.apply(c)
	}
}

type optionsPack []Option

func (o optionsPack) apply(c *config) {
	c.apply(o...)
}

// Skip returns an Option that adjusts the starting position of the retrieved stack.
func Skip(n int) Option {
	return skip(n)
}

type skip int

func (n skip) apply(c *config) {
	c.Skip += int(n)
}

// Limit returns an Option that limits the depth of the retrieved stack.
func Limit(n int) Option {
	return limit(n)
}

type limit int

func (n limit) apply(c *config) {
	c.Limit = int(n)
}

// Single is an Option that cancels Always and reverts to default behavior.
var Single single

type single struct{}

func (single) apply(c *config) {
	c.Always = false
}

// Always is an Option that specifies always including a new *Error in the error chain,
// even if the chain already contains a *Error.
//
// The default behavior of Wrap is to return the cause if the chain already contains a *Error,
// and to create a new *Error with the cause and stack frames otherwise.
//
// When this option is set, Wrap always creates a new *Error with the cause and stack frames.
//
// see
//
//	stacktrace.Always.Errorf(...)
var Always always

type always struct{}

func (always) apply(c *config) {
	c.Always = true
}

// Errorf is a syntax sugar for stacktrace.With(fmt.Errorf(format, a...), stacktrace.Always).
func (always) Errorf(format string, a ...any) error {
	return With(fmt.Errorf(format, a...), Skip(1), Always)
}

package internal

import (
	"time"
)

// Info contains additional information about the error.
type Info struct {
	// Suggestions is a list of suggestions for the user.
	Suggestions []string

	// Timestamp is the timestamp of the error.
	Timestamp time.Time

	// Context is the context of the error.
	Context map[string]any

	// StackTrace is the stack trace of the error.
	StackTrace []string

	// Inner is the inner error of the error.
	Inner error
}

// IsNil implements the errors.Pointer interface.
func (info *Info) IsNil() bool {
	return info == nil
}

// NewInfo creates a new Info.
//
// Returns:
//   - *Info: A pointer to the new Info. Never returns nil.
func NewInfo() *Info {
	return &Info{
		Suggestions: nil,
		Timestamp:   time.Now(),
		Context:     nil,
		StackTrace:  make([]string, 0),
		Inner:       nil,
	}
}

// Copy creates a shallow copy of the Info.
//
// Returns:
//   - *Info: A pointer to the new Info. Never returns nil.
func (info *Info) Copy() *Info {
	if info == nil {
		return NewInfo()
	}

	suggestions := make([]string, len(info.Suggestions))
	copy(suggestions, info.Suggestions)

	var context map[string]any

	if info.Context == nil {
		context = make(map[string]any)
	} else {
		context = make(map[string]any, len(info.Context))

		for key, value := range info.Context {
			context[key] = value
		}
	}

	var stack_trace []string

	if info.StackTrace == nil {
		stack_trace = make([]string, 0)
	} else {
		stack_trace = make([]string, len(info.StackTrace))
		copy(stack_trace, info.StackTrace)
	}

	return &Info{
		Suggestions: suggestions,
		Timestamp:   info.Timestamp,
		Context:     context,
		StackTrace:  stack_trace,
		Inner:       info.Inner,
	}
}

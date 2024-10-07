package errors

import (
	"fmt"

	"github.com/PlayerR9/go-errors/internal"
)

// Err represents a generalized error.
type Err struct {
	// Severity is the severity level of the error.
	Severity SeverityLevel

	// Code is the error code.
	Code ErrorCoder

	// Message is the error message.
	Message string

	*internal.Info
}

// Error implements the error interface.
func (e *Err) Error() string {
	if e == nil {
		return ""
	}

	var msg string

	if e.Message == "" {
		msg = "[no message was provided]"
	} else {
		msg = e.Message
	}

	return fmt.Sprintf("[%v] %v: %s", e.Severity, e.Code, msg)
}

// IsNil implements the Pointer interface.
func (e *Err) IsNil() bool {
	return e == nil
}

// New creates a new error.
//
// Parameters:
//   - code: The error code.
//   - message: The error message.
//
// Returns:
//   - *Err: A pointer to the new error. Never returns nil.
func New[C ErrorCoder](code C, message string) *Err {
	return &Err{
		Severity: ERROR,
		Code:     code,
		Message:  message,
		Info:     internal.NewInfo(),
	}
}

// NewWithSeverity creates a new error.
//
// Parameters:
//   - severity: The severity level of the error.
//   - code: The error code.
//   - message: The error message.
//
// Returns:
//   - *Err: A pointer to the new error. Never returns nil.
func NewWithSeverity[C ErrorCoder](severity SeverityLevel, code C, message string) *Err {
	return &Err{
		Severity: severity,
		Code:     code,
		Message:  message,
		Info:     internal.NewInfo(),
	}
}

// ChangeSeverity changes the severity level of the error. Does
// nothing if the receiver is nil.
//
// Parameters:
//   - new_severity: The new severity level of the error.
func (e *Err) ChangeSeverity(new_severity SeverityLevel) {
	if e == nil {
		return
	}

	e.Severity = new_severity
}

// AddSuggestion adds a suggestion to the error. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - suggestion: The suggestion to add.
func (e *Err) AddSuggestion(suggestion string) {
	if e == nil {
		return
	}

	// e.Suggestions = append(e.Suggestions, suggestion)
}

// AddContext adds a context to the error. Does nothing if the
// receiver is nil.
//
// Parameters:
//   - key: The key of the context.
//   - value: The value of the context.
func (e *Err) AddContext(key string, value any) {
	if e == nil {
		return
	}

	// if e.Context == nil {
	// 	e.Context = make(map[string]any)
	// }

	// e.Context[key] = value
}

/*
// Value returns the value of the context with the given key.
//
// Parameters:
//   - key: The key of the context.
//
// Returns:
//   - any: The value of the context with the given key.
//   - bool: true if the context contains the key, false otherwise.
func (e Err) Value(key string) (any, bool) {
	// if len(e.Context) == 0 {
	// 	return nil, false
	// }

	// value, ok := e.Context[key]
	// return value, ok
} */

// AddFrame prepends a frame to the stack trace. Does nothing
// if the receiver is nil or the trace is empty.
//
// Parameters:
//   - frame: The frame to add.
//
// If prefix is empty, the call is used as the frame. Otherwise a dot is
// added between the prefix and the call.
func (e *Err) AddFrame(frame string) {
	if e == nil || frame == "" {
		return
	}

	// e.StackTrace = append(e.StackTrace, frame)
}

// SetInner sets the inner error. Does nothing if the receiver is nil.
//
// Parameters:
//   - inner: The inner error.
func (e *Err) SetInner(inner error) {
	if e == nil {
		return
	}

	// e.Inner = inner
}

// NewFromError creates a new error from an error.
//
// Parameters:
//   - code: The error code.
//   - err: The error to wrap.
//
// Returns:
//   - *Err: A pointer to the new error. Never returns nil.
func NewFromError[C ErrorCoder](code C, err error) *Err {
	var outer *Err

	if err == nil {
		outer = &Err{
			Code:    code,
			Message: "something went wrong",
			Info:    internal.NewInfo(),
		}
	} else {
		switch inner := err.(type) {
		case *Err:
			outer = &Err{
				Code:    code,
				Message: inner.Message,
				Info:    inner.Info.Copy(),
			}

			inner.Info = nil // Clear any info since it is now in the outer error.
		default:
			outer = &Err{
				Code:    code,
				Message: inner.Error(),
				Info:    internal.NewInfo(),
			}
		}
	}

	outer.Severity = ERROR

	return outer
}

///////////////////////////////////////////////////////

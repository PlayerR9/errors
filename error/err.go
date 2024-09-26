package error

import (
	"fmt"
	"io"

	"github.com/PlayerR9/go-errors/error/internal"
)

// Err represents a generalized error.
type Err struct {
	// Severity is the severity level of the error.
	Severity SeverityLevel

	// Code is the error code.
	Code ErrorCoder

	// Message is the error message.
	Message string

	*Info
}

// Error implements the error interface.
func (e Err) Error() string {
	var msg string

	if e.Message == "" {
		msg = "[no message was provided]"
	} else {
		msg = e.Message
	}

	return fmt.Sprintf("[%v] %v: %s", e.Severity, e.Code, msg)
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
		Info:     NewInfo(),
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
		Info:     NewInfo(),
	}
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
			Info:    NewInfo(),
		}
	} else {
		switch inner := err.(type) {
		case *Err:
			// TODO: Handle this case.

			outer = &Err{
				Code:    code,
				Message: inner.Message,
				Info:    inner.Info,
			}

			inner.Info = nil // Clear any info since it is now in the outer error.
		default:
			// TODO: Handle this case.

			outer = &Err{
				Code:    code,
				Message: inner.Error(),
				Info:    NewInfo(),
			}
		}
	}

	outer.Severity = ERROR

	return outer
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

	e.Suggestions = append(e.Suggestions, suggestion)
}

// AddFrame prepends a frame to the stack trace. Does nothing
// if the receiver is nil or the trace is empty.
//
// Parameters:
//   - frame: The frame to add.
//
// If prefix is empty, the call is used as the frame. Otherwise a dot is
// added between the prefix and the call.
func (e *Err) AddFrame(frame string) {
	if e == nil {
		return
	}

	if e.StackTrace == nil {
		e.StackTrace = internal.NewStackTrace(frame)
	} else {
		e.StackTrace.AddFrame(frame)
	}
}

// SetInner sets the inner error. Does nothing if the receiver is nil.
//
// Parameters:
//   - inner: The inner error.
func (e *Err) SetInner(inner error) {
	if e == nil {
		return
	}

	e.Inner = inner
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

	if e.Context == nil {
		e.Context = make(map[string]any)
	}

	e.Context[key] = value
}

// Value returns the value of the context with the given key.
//
// Parameters:
//   - key: The key of the context.
//
// Returns:
//   - any: The value of the context with the given key.
//   - bool: true if the context contains the key, false otherwise.
func (e Err) Value(key string) (any, bool) {
	if len(e.Context) == 0 {
		return nil, false
	}

	value, ok := e.Context[key]
	return value, ok
}

// DisplayError displays the complete error to the writer.
//
// Parameters:
//   - w: The writer to write to.
//   - err: The error to display.
//
// Returns:
//   - error: The error that occurred while displaying the error.
func DisplayError(w io.Writer, err error) error {
	if err == nil {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(err.Error())

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	e, ok := err.(*Err)
	if !ok || e.Info == nil {
		return nil
	}

	err = e.DisplayInfo(w)
	if err != nil {
		return err
	}

	return nil
}

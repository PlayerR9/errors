package error

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/PlayerR9/go-errors/error/internal"
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
	StackTrace *internal.StackTrace

	// Inner is the inner error of the error.
	Inner error
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
		StackTrace:  nil,
		Inner:       nil,
	}
}

// Copy creates a shallow copy of the Info.
//
// Returns:
//   - *Info: A pointer to the new Info. Never returns nil.
func (info Info) Copy() *Info {
	return &Info{
		Suggestions: info.Suggestions,
		Timestamp:   info.Timestamp,
		Context:     info.Context,
		StackTrace:  info.StackTrace,
		Inner:       info.Inner,
	}
}

// IsNil checks whether the info is nil.
//
// Returns:
//   - bool: True if the info is nil, false otherwise.
func (info *Info) IsNil() bool {
	return info == nil
}

// DisplayInfo displays the info to the writer.
//
// Parameters:
//   - w: The writer to write to.
//
// Returns:
//   - error: The error that occurred while displaying the info.
func (info Info) DisplayInfo(w io.Writer) error {
	var b bytes.Buffer

	if !info.Timestamp.IsZero() {
		fmt.Fprintf(&b, "Occurred at: %v\n", info.Timestamp)
	}

	if len(info.Suggestions) > 0 {
		fmt.Fprintf(&b, "Suggestion: \n")

		for _, suggestion := range info.Suggestions {
			fmt.Fprintf(&b, "- %s\n", suggestion)
		}
	}

	if len(info.Context) > 0 {
		b.WriteString("\nContext:\n")

		for k, v := range info.Context {
			fmt.Fprintf(&b, "- %s: %v\n", k, v)
		}
	}

	if info.StackTrace != nil {
		fmt.Fprintf(&b, "\nStack trace:\n")
		fmt.Fprintf(&b, "- %s\n", info.StackTrace.String())
	}

	if info.Inner != nil {
		fmt.Fprintf(&b, "\nCaused by:\n")

		err := DisplayError(&b, info.Inner)
		if err != nil {
			return err
		}
	}

	data := b.Bytes()

	if len(data) == 0 {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}

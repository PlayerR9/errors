package errors

import (
	"fmt"
	"io"
	"strings"
)

// ZeroOf returns a zero value of type T.
//
// This is useful when you want to compare a value with its zero value.
//
// Example:
//
//	if x == (errors.ZeroOf[struct{}]()) {
//	    // x is the zero value of struct{}
//	}
func ZeroOf[T any]() T {
	return *new(T)
}

func WriteString(w io.Writer, str string) error {
	if str == "" {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(str)

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}

func DisplayError(w io.Writer, err error) error {
	if w == nil || err == nil {
		return nil
	}

	data := []byte(err.Error())

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	e, ok := As(err)
	if !ok {
		return nil
	}

	var builder strings.Builder

	if !e.Timestamp.IsZero() {
		fmt.Fprintf(&builder, "\nOccurred at: %v", e.Timestamp)
	}

	if len(e.Suggestions) > 0 {
		fmt.Fprintf(&builder, "\n\nSuggestion: ")

		for _, suggestion := range e.Suggestions {
			fmt.Fprintf(&builder, "\n- %s", suggestion)
		}
	}

	if len(e.Context) > 0 {
		fmt.Fprintf(&builder, "\n\nContext: ")

		for k, v := range e.Context {
			fmt.Fprintf(&builder, "\n- %s: %v", k, v)
		}
	}

	if e.StackTrace != nil {
		builder.WriteString("\nstack trace:\n\t")
		builder.WriteString(e.StackTrace.String())
	}

	if e.Inner != nil {
		builder.WriteString("\n\ncaused by: ")
		builder.WriteString(e.Inner.Error())
	}
}

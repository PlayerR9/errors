package errors

import (
	"io"

	gerr "github.com/PlayerR9/go-errors/error"
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

// DisplayError displays the complete error to the writer.
//
// Parameters:
//   - w: The writer to write to.
//   - err: The error to display.
//
// Returns:
//   - error: The error that occurred while displaying the error.
func DisplayError(w io.Writer, err error) error {
	return gerr.DisplayError(w, err)
}

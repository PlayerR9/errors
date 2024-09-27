package errors

import (
	"io"

	gerr "github.com/PlayerR9/go-errors/error"
)

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

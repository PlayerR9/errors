package errors

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/PlayerR9/go-errors/internal"
)

// display_info displays the info to the writer.
//
// Parameters:
//   - info: The info to display.
//   - w: The writer to write to.
//
// Returns:
//   - error: The error that occurred while displaying the info.
func display_info(info *internal.Info, w io.Writer) error {
	if info == nil {
		return nil
	}

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

		elem := make([]string, len(info.StackTrace))
		copy(elem, info.StackTrace)

		slices.Reverse(elem)

		fmt.Fprintf(&b, "- %s\n", strings.Join(elem, " <- "))
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

// Panic is like DisplayError but panics afterwards.
//
// Parameters:
//   - w: The writer to write to.
//   - to_display: The error to display.
func Panic(w io.Writer, to_display error) {
	if to_display == nil {
		return
	}

	e, ok := to_display.(*Err)
	if ok && e.Info != nil {
		err := display_info(e.Info, w)
		if err != nil {
			panic(err)
		}
	}

	panic(to_display)
}

// DisplayError displays the complete error to the writer.
//
// Parameters:
//   - w: The writer to write to.
//   - to_display: The error to display.
//
// Returns:
//   - error: The error that occurred while displaying the error.
func DisplayError(w io.Writer, to_display error) error {
	if to_display == nil {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(to_display.Error())

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	e, ok := to_display.(*Err)
	if !ok || e.Info == nil {
		return nil
	}

	err = display_info(e.Info, w)
	if err != nil {
		return err
	}

	return nil
}

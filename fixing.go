package errors

import (
	"strconv"

	gerr "github.com/PlayerR9/go-errors/error"
)

// Fixer is defines the behavior of an object that can be fixed. This
// must not have a non-pointer receiver.
type Fixer interface {
	// Fix fixes the object.
	//
	// Returns:
	//   - error: An error that occurred while fixing the object.
	Fix() error
}

// Fix fixes the object.
//
// Parameters:
//   - name: The name of the object.
//   - obj: The object to fix.
//   - allow_nil: Whether to allow the object to be nil.
//
// Returns:
//   - error: An error that occurred while fixing the object.
func Fix(name string, obj Fixer, allow_nil bool) error {
	if name == "" {
		name = "struct{}"
	}

	if obj == nil && !allow_nil {
		msg := strconv.Quote(name) + " must not be nil"

		err := gerr.New(FailFix, msg)
		return err
	}

	err := obj.Fix()
	if err == nil {
		return nil
	}

	new_err := gerr.NewFromError(FailFix, err)
	new_err.AddFrame(name + ".Fix()")

	return new_err
}

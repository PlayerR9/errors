package assert

import (
	"fmt"
	"os"
	"strconv"

	gers "github.com/PlayerR9/go-errors"
)

// Validater is an interface that all validaters must implement.
type Validater interface {
	// Validate validates the object.
	//
	// Returns:
	//   - error: An error that occurred while validating the object.
	Validate() error
}

// Validate validates the object.
//
// Parameters:
//   - name: The name of the object.
//   - obj: The object to validate.
//   - allow_nil: Whether to allow the object to be nil.
//
// Panics if the object's internal state is invalid.
func Validate(name string, obj Validater, allow_nil bool) {
	var err *gers.Err

	if obj == nil {
		if allow_nil {
			return
		}

		var msg string

		if name == "" {
			msg = "receiver must not be nil"
		} else {
			msg = strconv.Quote(name) + " must not be nil"
		}

		err = gers.NewWithSeverity(gers.FATAL, InvalidState, msg)
	} else {
		inner := obj.Validate()
		if inner == nil {
			return
		}

		err = gers.NewFromError(InvalidState, inner)
		err.ChangeSeverity(gers.FATAL)
	}

	var frame string

	if name == "" {
		frame = fmt.Sprintf("Validate[receiver, %t]", allow_nil)
	} else {
		frame = fmt.Sprintf("Validate[%q, %t]", name, allow_nil)
	}

	err.AddFrame(frame)

	gers.Panic(os.Stderr, err)
}

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
// Panics if the object can not be fixed.
func Fix(name string, obj Fixer, allow_nil bool) {
	var err *gers.Err

	if obj == nil {
		if allow_nil {
			return
		}

		var msg string

		if name == "" {
			msg = "receiver must not be nil"
		} else {
			msg = strconv.Quote(name) + " must not be nil"
		}

		err = gers.NewWithSeverity(gers.FATAL, FailFix, msg)
	} else {
		inner := obj.Fix()
		if inner == nil {
			return
		}

		err = gers.NewFromError(FailFix, inner)
		err.ChangeSeverity(gers.FATAL)
	}

	var frame string

	if name == "" {
		frame = fmt.Sprintf("Fix[receiver, %t]", allow_nil)
	} else {
		frame = fmt.Sprintf("Fix[%q, %t]", name, allow_nil)
	}

	err.AddFrame(frame)

	gers.Panic(os.Stderr, err)
}

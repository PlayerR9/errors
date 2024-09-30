package errors

import "fmt"

// ErrorCoder is an interface that all error codes must implement.
type ErrorCoder interface {
	// Int returns the integer value of the error code.
	//
	// Returns:
	//   - int: The integer value of the error code.
	Int() int

	fmt.Stringer
}

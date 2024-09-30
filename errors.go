package errors

import (
	"strconv"
)

// ErrorCode is the type of the error code.
type ErrorCode int

const (
	// BadParameter occurs when a parameter is invalid or is not
	// valid for some reason. For example, a nil pointer when nil
	// pointers are not allowed.
	BadParameter ErrorCode = iota

	// InvalidUsage occurs when users call a function without
	// proper setups or preconditions.
	InvalidUsage

	// NoSuchKey occurs when a context key is requested but does
	// not exist.
	NoSuchKey

	// OperationFail occurs when an operation cannot be completed
	// due to an internal error.
	OperationFail
)

// Int implements the error.ErrorCoder interface.
func (e ErrorCode) Int() int {
	return int(e)
}

// NewErrNilReceiver creates a new error.Err error with the code
// OperationFail.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilReceiver() *Err {
	err := New(OperationFail, "receiver must not be nil")
	err.AddSuggestion("Did you forget to initialize the receiver?")

	return err
}

// NewErrInvalidParameter creates a new error.Err error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
//
// This function is mostly useless since it just wraps BadParameter.
func NewErrInvalidParameter(message string) *Err {
	err := New(BadParameter, message)

	return err
}

// NewErrNilParameter creates a new error.Err error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilParameter(parameter string) *Err {
	msg := "parameter (" + strconv.Quote(parameter) + ") must not be nil"

	err := New(BadParameter, msg)
	err.AddSuggestion("Maybe you forgot to initialize the parameter?")

	return err
}

// NewErrInvalidUsage creates a new error.Err error.
//
// Parameters:
//   - message: The message of the error.
//   - usage: The usage/suggestion to solve the problem.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrInvalidUsage(message string, usage string) *Err {
	err := New(InvalidUsage, message)

	err.AddSuggestion(usage)

	return err
}

// NewErrNoSuchKey creates a new error.Err error.
//
// Parameters:
//   - key: The key that does not exist.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNoSuchKey(key string) *Err {
	err := New(NoSuchKey, "key ("+strconv.Quote(key)+") does not exist")

	return err
}

////////////////////////////////////////////////////////////

// NewErrAt creates a new error.Err error.
//
// Parameters:
//   - at: The operation at which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrAt(at string, reason error) *Err {
	var msg string

	if at == "" {
		msg = "an error occurred somewhere"
	} else {
		msg = "an error occurred at " + at
	}

	err := New(OperationFail, msg)
	err.SetInner(reason)

	return err
}

// NewErrAfter creates a new error.Err error.
//
// Parameters:
//   - before: The operation after which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrAfter(before string, reason error) *Err {
	var msg string

	if before == "" {
		msg = "an error occurred after something"
	} else {
		msg = "an error occurred after " + before
	}

	err := New(OperationFail, msg)
	err.SetInner(reason)

	return err
}

// NewErrBefore creates a new error.Err error.
//
// Parameters:
//   - after: The operation before which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrBefore(after string, reason error) *Err {
	var msg string

	if after == "" {
		msg = "an error occurred before something"
	} else {
		msg = "an error occurred before " + after
	}

	err := New(OperationFail, msg)
	err.SetInner(reason)

	return err
}

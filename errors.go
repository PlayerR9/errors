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
// Parameters:
//   - frame: The frame of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilReceiver(frame string) *Err {
	err := New(OperationFail, "receiver must not be nil")
	err.AddSuggestion("Did you forget to initialize the receiver?")

	err.AddFrame(frame)

	return err
}

// NewErrInvalidParameter creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - message: The message of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
//
// This function is mostly useless since it just wraps BadParameter.
func NewErrInvalidParameter(frame, message string) *Err {
	err := New(BadParameter, message)

	err.AddFrame(frame)

	return err
}

// NewErrNilParameter creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilParameter(frame, parameter string) *Err {
	msg := "parameter (" + strconv.Quote(parameter) + ") must not be nil"

	err := New(BadParameter, msg)
	err.AddSuggestion("Maybe you forgot to initialize the parameter?")

	return err
}

// NewErrInvalidUsage creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - message: The message of the error.
//   - usage: The usage/suggestion to solve the problem.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrInvalidUsage(frame, message, usage string) *Err {
	err := New(InvalidUsage, message)

	err.AddSuggestion(usage)

	err.AddFrame(frame)

	return err
}

// NewErrNoSuchKey creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - key: The key that does not exist.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNoSuchKey(frame, key string) *Err {
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

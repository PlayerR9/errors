package errors

import (
	"strconv"

	gerr "github.com/PlayerR9/go-errors/error"
)

//go:generate stringer -type=ErrorCode

type ErrorCode int

const (
	// BadParameter occurs when a parameter is invalid or is not
	// valid for some reason. For example, a nil pointer when nil
	// pointers are not allowed.
	BadParameter ErrorCode = iota

	// InvalidUsage occurs when users call a function without
	// proper setups or preconditions.
	InvalidUsage

	// FailFix occurs when a struct cannot be fixed or resolved
	// due to an invalid internal state.
	FailFix

	// OperationFail occurs when an operation cannot be completed
	// due to an internal error.
	OperationFail

	// NoSuchKey occurs when a context key is requested but does
	// not exist.
	NoSuchKey

	// AssertFail occurs when an assertion fails.
	AssertFail
)

// Int implements the error.ErrorCoder interface.
func (e ErrorCode) Int() int {
	return int(e)
}

// NewErrInvalidParameter creates a new error.Err error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrInvalidParameter(message string) *gerr.Err {
	err := gerr.New(BadParameter, message)

	return err
}

// NewErrNilParameter creates a new error.Err error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilParameter(parameter string) *gerr.Err {
	msg := "parameter (" + strconv.Quote(parameter) + ") must not be nil"

	err := gerr.New(BadParameter, msg)

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
func NewErrInvalidUsage(message string, usage string) *gerr.Err {
	err := gerr.New(InvalidUsage, message)

	err.AddSuggestion(usage)

	return err
}

// NewErrAt creates a new error.Err error.
//
// Parameters:
//   - at: The operation at which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrAt(at string, reason error) *gerr.Err {
	var msg string

	if at == "" {
		msg = "an error occurred somewhere"
	} else {
		msg = "an error occurred at " + at
	}

	err := gerr.New(OperationFail, msg)
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
func NewErrAfter(before string, reason error) *gerr.Err {
	var msg string

	if before == "" {
		msg = "an error occurred after something"
	} else {
		msg = "an error occurred after " + before
	}

	err := gerr.New(OperationFail, msg)
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
func NewErrBefore(after string, reason error) *gerr.Err {
	var msg string

	if after == "" {
		msg = "an error occurred before something"
	} else {
		msg = "an error occurred before " + after
	}

	err := gerr.New(OperationFail, msg)
	err.SetInner(reason)

	return err
}

// NewErrNoSuchKey creates a new error.Err error.
//
// Parameters:
//   - key: The key that does not exist.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNoSuchKey(key string) *gerr.Err {
	err := gerr.New(NoSuchKey, "key ("+strconv.Quote(key)+") does not exist")

	return err
}

// NewErrAssertFail creates a new error.Err error.
//
// Parameters:
//   - msg: The message of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrAssertFail(msg string) *gerr.Err {
	return gerr.NewWithSeverity(gerr.FATAL, AssertFail, msg)
}

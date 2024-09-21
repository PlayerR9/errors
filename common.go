package errors

import (
	"fmt"

	gcers "github.com/PlayerR9/errors/error"
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
)

// NewErrInvalidParameter creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrInvalidParameter(message string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, BadParameter, message)

	return err
}

// NewErrNilParameter creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrNilParameter(parameter string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, BadParameter, fmt.Sprintf("parameter (%q) cannot be nil", parameter))

	return err
}

// NewErrInvalidUsage creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//   - usage: The usage/suggestion to solve the problem.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrInvalidUsage(message string, usage string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, InvalidUsage, message)

	err.AddSuggestion(usage)

	return err
}

// NewErrFix creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrFix(message string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, FailFix, message)

	return err
}

package errors

import (
	gcers "github.com/PlayerR9/errors/error"
)

//go:generate stringer -type=ErrorCode

type ErrorCode int

const (
	BadParameter ErrorCode = iota
	InvalidUsage
	FailFix
)

// NewErrInvalidParameter creates a new ErrInvalidParameter error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//   - reason: the reason for the error.
//
// Returns:
//   - *ErrInvalidParameter: the new error. Never returns nil.
func NewErrInvalidParameter(format string, args ...any) *gcers.Err[ErrorCode] {
	err := gcers.NewErrF(gcers.FATAL, BadParameter, format, args...)

	return err
}

// NewErrNilParameter creates a new ErrInvalidParameter error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *ErrInvalidParameter: the new error. Never returns nil.
func NewErrNilParameter(parameter string) *gcers.Err[ErrorCode] {
	err := gcers.NewErrF(gcers.FATAL, BadParameter, "parameter %s cannot be nil", parameter)

	return err
}

// NewErrInvalidUsage creates a new ErrInvalidUsage error.
//
// Parameters:
//   - reason: The reason for the invalid usage.
//   - usage: The usage of the function.
//
// Returns:
//   - *ErrInvalidUsage: A pointer to the new ErrInvalidUsage error.
func NewErrInvalidUsage(message string, usage string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, InvalidUsage, message)

	err.AddSuggestion(usage)

	return err
}

// NewErrFix creates a new ErrFix error.
//
// Parameters:
//   - name: the name of the object.
//   - reason: the reason for the error.
//
// Returns:
//   - *ErrFix: the new error. Never returns nil.
func NewErrFix(message string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, FailFix, message)

	return err
}

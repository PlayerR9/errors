package errors

import (
	"errors"
	"fmt"

	gcers "github.com/PlayerR9/errors/error"
)

// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T gcers.ErrorCoder](err error, code T) bool {
	if err == nil {
		return false
	}

	var sub_err *gcers.Err

	ok := errors.As(err, &sub_err)
	if !ok {
		return false
	}

	other, ok := sub_err.Code.(T)
	return ok && other.Int() == code.Int()
}

// As returns the error if it is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - *gcers.Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func As(err error) (*gcers.Err, bool) {
	if err == nil {
		return nil, false
	}

	var sub_err *gcers.Err

	ok := errors.As(err, &sub_err)
	if !ok {
		return nil, false
	}

	return sub_err, true
}

// AsWithCode returns the error if it is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - *gcers.Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func AsWithCode[T gcers.ErrorCoder](err error, code T) (*gcers.Err, bool) {
	if err == nil {
		return nil, false
	}

	var sub_err *gcers.Err

	ok := errors.As(err, &sub_err)
	if !ok {
		return nil, false
	}

	other, ok := sub_err.Code.(T)
	if !ok || other.Int() != code.Int() {
		return nil, false
	}

	return sub_err, true
}

// Value is a function that returns the value of the context with the given key.
//
// Parameters:
//   - e: The error to get the value from.
//   - key: The key of the context.
//
// Returns:
//   - T: The value of the context with the given key.
//   - error: The error that occurred while getting the value.
func Value[C gcers.ErrorCoder, T any](e *gcers.Err, key string) (T, error) {
	if e == nil || len(e.Context) == 0 {
		return *new(T), NewErrNoSuchKey(key)
	}

	x, ok := e.Context[key]
	if !ok {
		return *new(T), NewErrNoSuchKey(key)
	}

	if x == nil {
		err := NewErrNoSuchKey(key)
		err.AddSuggestion("Found a key with the same name but has a nil value")

		return *new(T), err
	}

	val, ok := x.(T)
	if !ok {
		err := NewErrNoSuchKey(key)
		err.AddSuggestion(fmt.Sprintf("Found a key with the same name but has a value of type %T", x))

		return *new(T), err
	}

	return val, nil
}

/*
// LimitErrorMsg is a function that limits the number of errors in an error chain.
//
// Parameters:
//   - err: The error to limit.
//   - limit: The maximum number of errors to limit.
//
// Returns:
//   - error: The limited error.
//
// If the error is nil or the limit is less than or equal to 0, the function returns nil.
func LimitErrorMsg(err error, limit int) error {
	if err == nil || limit <= 0 {
		return nil
	}

	target := err

	for i := 0; i < limit; i++ {
		w, ok := target.(Unwrapper)
		if !ok {
			return err
		}

		reason := w.Unwrap()
		if reason == nil {
			return err
		}

		target = reason
	}

	if target == nil {
		return err
	}

	w, ok := target.(Unwrapper)
	if !ok {
		return err
	}

	w.ChangeReason(nil)

	return err
} */

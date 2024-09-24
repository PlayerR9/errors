package errors

import "fmt"

// Assert asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and the given message.
//
// Parameters:
//   - cond: The condition to assert.
//   - msg: The message to use if the condition is false.
func Assert(cond bool, msg string) {
	if cond {
		return
	}

	panic(NewErrAssertFail(msg))
}

// AssertF asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the formatted string.
//
// Parameters:
//   - cond: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func AssertF(cond bool, format string, args ...any) {
	if cond {
		return
	}

	msg := fmt.Sprintf(format, args...)

	panic(NewErrAssertFail(msg))
}

// AssertErr asserts that an error is nil.
//
// If the error is not nil, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original error.
//
// Parameters:
//   - err: The error to check.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func AssertErr(err error, format string, args ...any) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = " + err.Error()

	panic(NewErrAssertFail(msg))
}

// AssertOk asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original condition.
//
// Parameters:
//   - ok: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func AssertOk(ok bool, format string, args ...any) {
	if ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = false"

	panic(NewErrAssertFail(msg))
}

// AssertNotOk asserts that a condition is false.
//
// If the condition is true, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original condition.
//
// Parameters:
//   - ok: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func AssertNotOk(ok bool, format string, args ...any) {
	if !ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = true"

	panic(NewErrAssertFail(msg))
}

// AssertNotNil asserts that the given object is not nil.
//
// If the object is nil, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object.
//
// Parameters:
//   - obj: The object to assert is not nil.
//   - name: The name of the object to use for the error message.
func AssertNotNil(obj any, name string) {
	if obj != nil {
		return
	}

	if name == "" {
		name = "object"
	}

	msg := name + " = nil"

	panic(NewErrAssertFail(msg))
}

// AssertNotZero asserts that the given object is not its zero value.
//
// If the object is its zero value, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object and its
// zero value.
//
// Parameters:
//   - obj: The object to assert is not its zero value.
//   - name: The name of the object to use for the error message.
func AssertNotZero[T comparable](obj T, name string) {
	zero := ZeroOf[T]()

	if obj != zero {
		return
	}

	if name == "" {
		name = "object"
	}

	msg := fmt.Sprintf("%s = %v", name, obj)

	panic(NewErrAssertFail(msg))
}

// AssertType asserts that the given object is of type T.
//
// If the object is not of type T, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object and its
// expected type.
//
// Parameters:
//   - obj: The object to assert is of type T.
//   - name: The name of the object to use for the error message.
//   - allow_nil: Whether to allow the object to be nil.
func AssertType[T any](obj any, name string, allow_nil bool) {
	if name == "" {
		name = "object"
	}

	zero := ZeroOf[T]()

	var msg string

	if obj == nil {
		msg = fmt.Sprintf("%s = nil, expected %T", name, zero)
	} else {
		_, ok := obj.(T)
		if ok {
			return
		}

		msg = fmt.Sprintf("%s = %T, expected %T", name, obj, zero)
	}

	panic(NewErrAssertFail(msg))
}

// AssertConv asserts that the given object can be converted to type T.
//
// If the object can be converted to type T, it returns the converted value.
// Otherwise, it calls panic() with an error.Err with the code AssertFail and a
// message that includes the original object and its expected type.
//
// Parameters:
//   - obj: The object to convert to type T.
//   - name: The name of the object to use for the error message.
//
// Returns:
//   - T: The converted value.
func AssertConv[T any](obj any, name string) T {
	if name == "" {
		name = "object"
	}

	zero := ZeroOf[T]()

	var msg string

	if obj == nil {
		msg = fmt.Sprintf("%s = nil, expected %T", name, zero)
	} else {
		val, ok := obj.(T)
		if ok {
			return val
		}

		msg = fmt.Sprintf("%s = %T, expected %T", name, obj, zero)
	}

	panic(NewErrAssertFail(msg))
}

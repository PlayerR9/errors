package assert

import (
	"fmt"
	"os"

	gers "github.com/PlayerR9/go-errors"
)

// Cond asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and the given message.
//
// Parameters:
//   - cond: The condition to assert.
//   - msg: The message to use if the condition is false.
func Cond(cond bool, msg string) {
	if cond {
		return
	}

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.Cond")

	gers.Panic(os.Stderr, err)
}

// CondF asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the formatted string.
//
// Parameters:
//   - cond: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func CondF(cond bool, format string, args ...any) {
	if cond {
		return
	}

	msg := fmt.Sprintf(format, args...)

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.CondF")

	gers.Panic(os.Stderr, err)
}

// Err asserts that an error is nil.
//
// If the error is not nil, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original error.
//
// Parameters:
//   - err: The error to check.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func Err(inner error, format string, args ...any) {
	if inner == nil {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = " + inner.Error()

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.Err")

	gers.Panic(os.Stderr, err)
}

// Ok asserts that a condition is true.
//
// If the condition is false, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original condition.
//
// Parameters:
//   - ok: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func Ok(ok bool, format string, args ...any) {
	if ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = false"

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.Ok")

	gers.Panic(os.Stderr, err)
}

// NotOk asserts that a condition is false.
//
// If the condition is true, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original condition.
//
// Parameters:
//   - ok: The condition to assert.
//   - format: The format string to use for the message.
//   - args: The arguments to pass to the format string.
func NotOk(ok bool, format string, args ...any) {
	if !ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = true"

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.NotOk")

	gers.Panic(os.Stderr, err)
}

// NotNil asserts that the given object is not nil.
//
// If the object is nil, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object.
//
// Parameters:
//   - obj: The object to assert is not nil.
//   - name: The name of the object to use for the error message.
func NotNil(obj any, name string) {
	if obj != nil {
		return
	}

	if name == "" {
		name = "object"
	}

	msg := name + " = nil"

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.NotNil")

	gers.Panic(os.Stderr, err)
}

// NotZero asserts that the given object is not its zero value.
//
// If the object is its zero value, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object and its
// zero value.
//
// Parameters:
//   - obj: The object to assert is not its zero value.
//   - name: The name of the object to use for the error message.
func NotZero[T comparable](obj T, name string) {
	zero := *new(T)

	if obj != zero {
		return
	}

	if name == "" {
		name = "object"
	}

	msg := fmt.Sprintf("%s = %v", name, obj)

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.NotZero")

	gers.Panic(os.Stderr, err)
}

// Type asserts that the given object is of type T.
//
// If the object is not of type T, it calls panic() with an error.Err with the
// code AssertFail and a message that includes the original object and its
// expected type.
//
// Parameters:
//   - obj: The object to assert is of type T.
//   - name: The name of the object to use for the error message.
//   - allow_nil: Whether to allow the object to be nil.
func Type[T any](obj any, name string, allow_nil bool) {
	if name == "" {
		name = "object"
	}

	zero := *new(T)

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

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.Type")

	gers.Panic(os.Stderr, err)
}

// Conv asserts that the given object can be converted to type T.
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
func Conv[T any](obj any, name string) T {
	if name == "" {
		name = "object"
	}

	zero := *new(T)

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

	err := gers.NewWithSeverity(gers.FATAL, AssertFail, msg)
	err.AddFrame("assert.Conv")

	gers.Panic(os.Stderr, err)

	panic("unreachable")
}

// New asserts a constructor returns a non-zero value.
//
// If the constructor returns a zero value, it calls panic() with an error.Err
// with the code AssertFail and a message that includes the original error.
//
// Parameters:
//   - obj: The object returned by the constructor.
//   - inner: The error returned by the constructor.
//
// Returns:
//   - T: The non-zero value returned by the constructor.
func New[T gers.Pointer](obj T, inner error) T {
	var err *gers.Err

	if inner != nil {
		err = gers.NewFromError(AssertFail, inner)
		err.ChangeSeverity(gers.FATAL)
	} else {
		if !obj.IsNil() {
			return obj
		}

		err = gers.NewWithSeverity(gers.FATAL, AssertFail, "object must not be the zero value")
	}

	err.AddFrame("assert.New")

	gers.Panic(os.Stderr, err)

	panic("unreachable")
}

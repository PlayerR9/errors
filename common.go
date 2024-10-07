package errors

import (
	"errors"
	"fmt"

	"github.com/PlayerR9/go-errors/internal"
)

// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T ErrorCoder](err error, code T) bool {
	if err == nil {
		return false
	}

	var sub_err *Err

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
//   - *Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func As(err error) (*Err, bool) {
	if err == nil {
		return nil, false
	}

	var sub_err *Err

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
//   - *Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func AsWithCode[T ErrorCoder](err error, code T) (*Err, bool) {
	if err == nil {
		return nil, false
	}

	var sub_err *Err

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

/*
// Value is a function that returns the value of the context with the given key.
//
// Parameters:
//   - e: The error to get the value from.
//   - key: The key of the context.
//
// Returns:
//   - T: The value of the context with the given key.
//   - error: The error that occurred while getting the value.
func Value[C ErrorCoder, T any](e *Err, key string) (T, error) {
	// zero := *new(T)

	// if e == nil || len(e.Context) == 0 {
	// 	return zero, NewErrNoSuchKey("Value()", key)
	// }

	// x, ok := e.Context[key]
	// if !ok {
	// 	return zero, NewErrNoSuchKey("Value()", key)
	// }

	// if x == nil {
	// 	err := NewErrNoSuchKey("Value()", key)
	// 	err.AddSuggestion("Found a key with the same name but has a nil value")

	// 	return zero, err
	// }

	// val, ok := x.(T)
	// if !ok {
	// 	err := NewErrNoSuchKey("Value()", key)
	// 	err.AddSuggestion(fmt.Sprintf("Found a key with the same name but has a value of type %T", x))

	// 	return zero, err
	// }

	// return val, nil
} */

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

// Merge merges the inner Info into the outer Info.
//
// Parameters:
//   - outer: The outer Info to merge.
//   - inner: The inner Info to merge.
//
// Returns:
//   - *Info: A pointer to the new Info. Never returns nil.
//
// Note:
//   - The other Info is the inner info of the current Info and, as such,
//     when conflicts occur, the outer Info takes precedence.
func Merge(outer, inner *internal.Info) *internal.Info {
	if inner == nil {
		return outer.Copy()
	}

	// suggestions := make([]string, 0, len(outer.Suggestions)+len(inner.Suggestions))
	// suggestions = append(suggestions, outer.Suggestions...)
	// suggestions = append(suggestions, inner.Suggestions...)

	// context := make(map[string]any)

	// for key, value := range inner.Context {
	// 	context[key] = value
	// }

	// for key, value := range outer.Context {
	// 	context[key] = value
	// }

	// stack_trace := make([]string, 0, len(outer.StackTrace)+len(inner.StackTrace))
	// stack_trace = append(stack_trace, outer.StackTrace...)
	// stack_trace = append(stack_trace, inner.StackTrace...)

	return &internal.Info{
		// Suggestions: suggestions,
		// Timestamp:   outer.Timestamp,
		// Context:    context,
		// StackTrace: stack_trace,
		// Inner: MergeErrors(outer.Inner, inner.Inner),
	}
}

func MergeErrors(outer, inner error) error {
	if outer == nil {
		return inner
	} else if inner == nil {
		return outer
	}

	o, ok1 := outer.(*Err)
	i, ok2 := inner.(*Err)

	if !ok1 && !ok2 {
		return fmt.Errorf("%w: %w", outer, inner)
	}

	var err *Err

	if ok1 {
		err = &Err{
			Severity: o.Severity,
			Code:     o.Code,
			Message:  o.Message,
		}
	} else {
		err = &Err{
			Severity: i.Severity,
			Code:     i.Code,
			Message:  i.Message,
		}
	}

	err.Info = Merge(o.Info, i.Info)

	// if ok1 && !ok2 {
	// 	err.Info.Inner = MergeErrors(o.Info.Inner, i)
	// } else if !ok1 && ok2 {
	// 	err.Info.Inner = MergeErrors(o, i.Info.Inner)
	// }

	return err
}

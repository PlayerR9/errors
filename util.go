package errors

// ZeroOf returns a zero value of type T.
//
// This is useful when you want to compare a value with its zero value.
//
// Example:
//
//	if x == (errors.ZeroOf[struct{}]()) {
//	    // x is the zero value of struct{}
//	}
func ZeroOf[T any]() T {
	return *new(T)
}

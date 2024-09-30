package errors

// Pointer is an interface that checks whether a pointer is nil.
type Pointer interface {
	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

package assert

// ErrorCode is the type of the error code.
type ErrorCode int

const (
	// AssertFail occurs when a test or assertion fails.
	AssertFail ErrorCode = iota

	// InvalidState is a type of assertion that occurs when
	// a method is called on a struct having an invalid state.
	InvalidState

	// FailFix occurs when a struct cannot be fixed or resolved
	// due to an invalid internal state.
	FailFix
)

// Int implements the error.ErrorCoder interface.
func (e ErrorCode) Int() int {
	return int(e)
}

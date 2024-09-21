// Code generated by "stringer -type=SeverityLevel"; DO NOT EDIT.

package error

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[INFO-0]
	_ = x[WARNING-1]
	_ = x[ERROR-2]
	_ = x[FATAL-3]
}

const _SeverityLevel_name = "INFOWARNINGERRORFATAL"

var _SeverityLevel_index = [...]uint8{0, 4, 11, 16, 21}

func (i SeverityLevel) String() string {
	if i < 0 || i >= SeverityLevel(len(_SeverityLevel_index)-1) {
		return "SeverityLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SeverityLevel_name[_SeverityLevel_index[i]:_SeverityLevel_index[i+1]]
}

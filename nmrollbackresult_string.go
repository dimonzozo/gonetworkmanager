// Code generated by "stringer -type=NMRollbackResult"; DO NOT EDIT.

package gonetworkmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NMRollbackResultOk-0]
	_ = x[NMRollbackResultErrNoDevice-1]
	_ = x[NMRollbackResultErrDeviceUnmanaged-2]
	_ = x[NMRollbackResultErrFailed-3]
}

const _NMRollbackResult_name = "NMRollbackResultOkNMRollbackResultErrNoDeviceNMRollbackResultErrDeviceUnmanagedNMRollbackResultErrFailed"

var _NMRollbackResult_index = [...]uint8{0, 18, 45, 79, 104}

func (i NMRollbackResult) String() string {
	if i >= NMRollbackResult(len(_NMRollbackResult_index)-1) {
		return "NMRollbackResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NMRollbackResult_name[_NMRollbackResult_index[i]:_NMRollbackResult_index[i+1]]
}

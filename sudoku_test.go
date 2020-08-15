package main

import "testing"

func TestChkRow(t *testing.T) {
	// TODO - set up a completed board to test for success, edit to test failures
	validRow := board{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	invalidRow := board{[]int{9, 2, 3, 4, 3, 6, 7, 8, 9}}
	zeroRow := board{[]int{0, 2, 3, 0, 0, 6, 7, 0, 9}}

	if !chkRow(0, &validRow) {
		t.Errorf("Row %d didn't pass as a valid row when it should", validRow[0])
	}
	if chkRow(0, &invalidRow) {
		t.Errorf("Row %d passed as a valid row when it shouldn't", invalidRow[0])
	}
	if !chkRow(0, &zeroRow) {
		t.Errorf("Row %d didn't pass as a valid row when it should", zeroRow[0])
	}
}

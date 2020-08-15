package main

import "testing"

func TestChkLine(t *testing.T) {
	// TODO - set up a completed board to test for success, edit to test failures
	validRow := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	invalidRow := []int{9, 2, 3, 4, 3, 6, 7, 8, 9}
	zeroRow := []int{0, 2, 3, 0, 0, 6, 7, 0, 9}

	if !chkLine(validRow) {
		t.Errorf("Row %d didn't pass as a valid row when it should", validRow[0])
	}
	if chkLine(invalidRow) {
		t.Errorf("Row %d passed as a valid row when it shouldn't", invalidRow[0])
	}
	if !chkLine(zeroRow) {
		t.Errorf("Row %d didn't pass as a valid row when it should", zeroRow[0])
	}
}

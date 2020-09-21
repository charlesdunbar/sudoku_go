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

func TestValidateBoard(t *testing.T) {
	// Gen a fully random board, check values match and row/column/box is expected
}

func TestGenBoard(t *testing.T) {
	b := genBoard("......9.....5....85..83....35..82.....6...2....937..4.76........3.4.5.76..2......")

	if b[0][6].value != 9 {
		t.Errorf("Board did not generate correctly, expected 9 at line 0 column 6, got %d", b[0][6].value)
	}
	if b[0][6].box != 3 {
		t.Errorf("Board did not generate correctly, expected line 0 colunn 6 to be in box 3, actually %d", b[0][6].box)
	}

}

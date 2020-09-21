package main

import (
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

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

func findBox(x, y int) int {
	switch x {
	case 0, 1, 2:
		switch y {
		case 0, 1, 2:
			return 1
		case 3, 4, 5:
			return 2
		case 6, 7, 8:
			return 3
		}
	case 3, 4, 5:
		switch y {
		case 0, 1, 2:
			return 4
		case 3, 4, 5:
			return 5
		case 6, 7, 8:
			return 6
		}
	case 6, 7, 8:
		switch y {
		case 0, 1, 2:
			return 7
		case 3, 4, 5:
			return 8
		case 6, 7, 8:
			return 9
		}
	}
	return 0
}

func TestGenBoard(t *testing.T) {

	properties := gopter.NewProperties(nil)

	properties.Property("should generate boards correctly", prop.ForAll(
		func(puzzle string) bool {
			b := genBoard(puzzle)
			stringCount := 0
			spl := strings.Split(puzzle, "")
			for row := 0; row < 9; row++ {
				for col := 0; col < 9; col++ {
					var val int
					if spl[stringCount] == "." {
						val = 0
					} else {
						var err error
						val, err = strconv.Atoi(spl[stringCount])
						if err != nil {
							log.Fatal(err)
						}
					}

					if !(b[row][col].x == row &&
						b[row][col].y == col &&
						b[row][col].value == val &&
						b[row][col].box == findBox(row, col)) {
						log.Fatalf("Cell %d %d from \n%vshould have x == %d, y == %d, box == %d, and value == %d, but actually was\n%+v\n",
							row, col, b, row, col, val, findBox(row, col), b[row][col])
						return false
					}
					stringCount++
				}
			}
			return true

		},
		gen.RegexMatch(`[0-9\.]{81}`),
	))
	properties.TestingRun(t)

}

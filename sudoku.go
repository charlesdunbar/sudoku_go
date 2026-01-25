package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// board represents a 9x9 Sudoku board as a slice of rows,
// where each row is a slice of `cell` values.
// Indexing uses `board[x][y]` where x is the row and y is the column.
type board [][]cell

// cell holds the coordinates and current value for a single
// Sudoku cell. Fields:
// - x, y: zero-based row and column indices
// - box: 1-based ID of the 3x3 sub-box the cell belongs to
// - value: the current value in the cell (0 means empty)
type cell struct {
	x, y, box, value int
}

// String returns a simple textual representation of the board
// showing the numeric values for each cell (0 for empty).
// This is intended for debugging and human-readable output.
func (b board) String() string {
	r := ""
	for i := range b {
		for j := range b[i] {
			r += fmt.Sprintf("%d ", b[i][j].value)
		}
		r += "\n"
	}
	return r
}

// genBoard creates a `board` from a 81-character puzzle string `p`.
// The input string should contain digits '1'..'9' for known cells
// and '.' for empty cells. The function populates each `cell` with
// its coordinates, which 3x3 box it belongs to, and initial value.
func genBoard(p string) board {
	b := make(board, 9)
	for x := 0; x < 9; x++ {
		b[x] = make([]cell, 9)
	}
	spl := strings.Split(p, "")

	count := 0
	boxID := 1
	for i := 0; i < 9; i++ {
		boxRowCounter := 1
		// Determine starting box ID based off how far into the row (i) loop we are
		switch {
		case i < 3:
			boxID = 1
		case i >= 3 && i <= 5:
			boxID = 4
		case i > 5:
			boxID = 7
		}
		for j := 0; j < 9; j++ {
			b[i][j].x = i
			b[i][j].y = j
			// Same boxID 3 times, then increment the ID for the next box
			if boxRowCounter == 4 {
				boxRowCounter = 1
				boxID++
			}
			b[i][j].box = boxID
			boxRowCounter++
			// Already at 0 since cell is an []int array, can continue
			if spl[count] == "." {
				count++
				continue
			} else {
				val, err := strconv.Atoi(spl[count])
				if err != nil {
					log.Fatal(err)
				}
				b[i][j].value = val
				count++
			}
		}
	}
	return b
}

func chkLine(x []int) bool {
	// Check that the non-zero elements in x contain no duplicates.
	row := removeZeroes(x)
	uniq := unique(row)
	return len(uniq) == len(row)
}

// validateBoard verifies that placing the current value at
// coordinates `(x,y)` does not violate Sudoku rules in its
// row, column, or 3x3 box. It returns true if all three checks pass.
//
// NOTE: the TODO in the original code suggests running the three
// checks concurrently; that's safe because the checks only read
// from the board and do not mutate shared state — they can be
// parallelized with goroutines and synchronized with channels or
// a `sync.WaitGroup` if desired.
func validateBoard(x, y int, b *board) bool {
	// TODO - These can be concurrent
	// OLD CODE:
	// genRow := lineFilter(b, func(val cell) bool {
	// 	return val.x == x
	// })
	// row := chkLine(genRow)

	// genCol := lineFilter(b, func(val cell) bool {
	// 	return val.y == y
	// })
	// col := chkLine(genCol)

	// genBox := lineFilter(b, func(val cell) bool {
	// 	return val.box == (*b)[x][y].box
	// })
	// box := chkLine(genBox)

	// return row && col && box
	// CONCURRENT CODE:
	// ch := make(chan bool, 3)
	// go func() { ch <- chkLine(lineFilter(b, func(c cell) bool { return c.x == x })) }()
	// go func() { ch <- chkLine(lineFilter(b, func(c cell) bool { return c.y == y })) }()
	// go func() { ch <- chkLine(lineFilter(b, func(c cell) bool { return c.box == (*b)[x][y].box })) }()
	// for range 3 {
	// 	if !<-ch {
	// 		return false
	// 	}
	// }
	// return true

	// INLINE CODE:
	// row
	var seen [10]bool
	for j := 0; j < 9; j++ {
		v := (*b)[x][j].value
		if v == 0 {
			continue
		}
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	// column
	seen = [10]bool{}
	for i := 0; i < 9; i++ {
		v := (*b)[i][y].value
		if v == 0 {
			continue
		}
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	// box (use index arithmetic, avoids box lookups/slices)
	sx := (x / 3) * 3
	sy := (y / 3) * 3
	seen = [10]bool{}
	for i := sx; i < sx+3; i++ {
		for j := sy; j < sy+3; j++ {
			v := (*b)[i][j].value
			if v == 0 {
				continue
			}
			if seen[v] {
				return false
			}
			seen[v] = true
		}
	}
	return true
}

// solve is a simple backtracking solver that fills zero (empty)
// cells in the board. It traverses the list of empty cells in
// order, incrementing cell values from 1..9 and backtracking when
// no valid candidate is found. The function prints progress and
// exits if it takes too many steps.
func solve(b *board) {
	stepCount := 0
	zeroBoard := cellFilter(b, func(val cell) bool {
		return val.value == 0
	})
	for i := 0; i < len(zeroBoard); i++ {
		stepCount++
		if stepCount%1000 == 0 {
			fmt.Printf("Current step: %d\n\n%v", stepCount, b)
		}
		if stepCount > 1000000 {
			fmt.Println("Took too long")
			os.Exit(42)
		}
		for j := 0; j < 10; j++ {
			zeroBoard[i].value++
			if zeroBoard[i].value == 10 {
				zeroBoard[i].value = 0
				i -= 2 // Go back 2 spaces to increment previous cell by one and try this cell 1-9 again
				break
			}

			if validateBoard(zeroBoard[i].x, zeroBoard[i].y, b) {
				break
			}
		}
	}
	fmt.Printf("Finished after %d steps\n", stepCount)
}

// cellFilter returns pointers to cells for which predicate `f` returns
// true. The returned slice contains pointers into the original board
// so mutating the returned `*cell` elements updates the board.
func cellFilter(b *board, f func(cell) bool) []*cell {
	line := []*cell{}
	for i := range *b {
		for j := range (*b)[i] {
			if f((*b)[i][j]) {
				line = append(line, &(*b)[i][j])
			}
		}
	}
	return line
}

// lineFilter returns the integer values of cells that satisfy the
// predicate `f`. It's used to extract a row, column, or box as a
// slice of ints for validation helpers such as `chkLine`.
func lineFilter(b *board, f func(cell) bool) []int {
	line := []int{}
	for i := range *b {
		for j := range (*b)[i] {
			if f((*b)[i][j]) {
				line = append(line, (*b)[i][j].value)
			}
		}
	}
	return line
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	p := loadCsv("./puzzles.csv")
	if len(p) == 0 {
		log.Fatal("no puzzles loaded")
	}
	// choose a random puzzle index and build the board from it
	puzzle := rand.Intn(len(p))
	b := genBoard(p[puzzle])
	fmt.Printf("Solving puzzle number %d from file:\n\n%v\n", puzzle+1, b)
	start := time.Now()
	solve(&b)
	elapsed := time.Since(start)
	fmt.Printf("Puzzle %d took %v to run\n\n", puzzle+1, elapsed)
	fmt.Println(b)
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type board [][]cell

type cell struct {
	x, y, value int
}

func (b board) String() string {
	r := ""
	for i := range b {
		for j := range b[i] {
			/*if b[i][j] == 0 {
				r += " "
			} else {
				r += fmt.Sprintf("%d ", b[i][j])
			}*/
			r += fmt.Sprintf("%d ", b[i][j].value)
		}
		r += "\n"
	}
	return r
}

func genBoard(p string) board {
	b := make(board, 9)
	for x := 0; x < 9; x++ {
		b[x] = make([]cell, 9)
	}
	spl := strings.Split(p, "")

	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[i][j].x = i
			b[i][j].y = j
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
	row := removeZeroes(x)
	uniq := unique(row)
	return len(uniq) == len(row)
}

/*func chkBox(x, y int, b *board) bool {
	// Generate an []int to pass to chkLine based on the 3x3 square x and y are in
	val := [][]int{[]int{x, y}}
	box := []int{}

	box1 := [][]int{[]int{0, 0}, []int{0, 1}, []int{0, 2}, []int{1, 0}, []int{1, 1}, []int{1, 2}, []int{2, 0}, []int{2, 1}, []int{2, 2}}

	//[0][0], [0][1], [0][2], [1][0], [1][1], [1][2], [2][0], [2][1], [2][2] = box 1


	return true
} /*

// From https://stackoverflow.com/a/10485970
/*func contains(s [][]int, e []int) bool {
	for i, j := range s {
		if s[i] == e {
			return true
		}
	}
	return false
}*/

// From https://www.golangprograms.com/remove-duplicate-values-from-slice.html
// Takes a slice and returns only unique values
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeZeroes(intSlice []int) []int {
	r := []int{}
	for i := range intSlice {
		if intSlice[i] != 0 {
			r = append(r, intSlice[i])
		}
	}
	return r
}

func validateBoard(x, y int, b *board) bool {
	// TODO - These can be concurrent
	genRow := lineFilter(b, func(val cell) bool {
		return val.x == x
	})
	row := chkLine(genRow)

	genCol := lineFilter(b, func(val cell) bool {
		return val.y == y
	})
	col := chkLine(genCol)

	//box := chkBox(x, y, b)
	return row && col
}

/*func solve(b *board) {
	var curRow, curCol int = 0, 0
	for {
		(*b)[curRow][curCol]++
		validateBoard(curRow, curCol, b)
	}

}*/

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
	b := genBoard("......9.....5....85..83....35..82.....6...2....937..4.76........3.4.5.76..2......")
	fmt.Println(b)
	fmt.Println(validateBoard(0, 0, &b))
}

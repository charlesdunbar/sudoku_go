package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type board [][]int

func (b board) String() string {
	r := ""
	for i := range b {
		for j := range b[i] {
			/*if b[i][j] == 0 {
				r += " "
			} else {
				r += fmt.Sprintf("%d ", b[i][j])
			}*/
			r += fmt.Sprintf("%d ", b[i][j])
		}
		r += "\n"
	}
	return r
}

func genBoard(p string) board {
	b := make(board, 9)
	for x := 0; x < 9; x++ {
		b[x] = make([]int, 9)
	}
	spl := strings.Split(p, "")

	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if spl[count] == "." {
				count++
				continue
			} else {
				val, err := strconv.Atoi(spl[count])
				if err != nil {
					log.Fatal(err)
				}
				b[i][j] = val
				count++
			}
		}
	}
	return b
}

func chkRow(x int, b *board) bool {
	row := removeZeroes((*b)[x])
	uniq := unique(row)
	return len(uniq) == len(row)
}

func chkCol(y int, b *board) bool {
	return true
}

func chkBox(b *board) bool {
	return true
}

// From https://www.golangprograms.com/remove-duplicate-values-from-slice.html
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
	row := chkRow(x, b)
	col := chkCol(y, b)
	box := chkBox(b)
	return row && col && box
}

func solve(b *board) {
	var curRow, curCol int = 0, 0
	for {
		(*b)[curRow][curCol]++
		validateBoard(curRow, curCol, b)
	}

}

func main() {
	b := genBoard("......9.....5....85..83....35..82.....6...2....937..4.76........3.4.5.76..2......")
	fmt.Println(b)
}

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

type board [][]cell

type cell struct {
	x, y, box, value int
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
	row := removeZeroes(x)
	uniq := unique(row)
	return len(uniq) == len(row)
}

func validateBoard(x, y int, b *board) bool {
	// TODO - These can be concurrent
	//fmt.Println("***Generating Row Filter***")
	genRow := lineFilter(b, func(val cell) bool {
		return val.x == x
	})
	row := chkLine(genRow)

	//fmt.Println("***Generating Column Filter***")
	genCol := lineFilter(b, func(val cell) bool {
		return val.y == y
	})
	col := chkLine(genCol)

	//fmt.Println("***Generating Box Filter***")
	genBox := lineFilter(b, func(val cell) bool {
		return val.box == (*b)[x][y].box
	})
	box := chkLine(genBox)

	return row && col && box
}

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
		//fmt.Printf("Start of loop - i is %d, stepCount is %d\n", i, stepCount)
		if stepCount > 1000000 {
			fmt.Printf("Took too long")
			os.Exit(42)
		}
		for j := 0; j < 9; j++ {
			zeroBoard[i].value++
			if zeroBoard[i].value == 10 {
				//fmt.Printf("Cell %+v is invalid, going back two columns, currently at %d\n", (*b)[zeroBoard[i].x][zeroBoard[i].y], i)
				zeroBoard[i].value = 0
				i -= 2
				//fmt.Printf("Current board is now\n%v", b)
				break
			}

			if !validateBoard(zeroBoard[i].x, zeroBoard[i].y, b) {
				//fmt.Printf("Cell %+v is invalid, going back a column to increment value, i is currently %d, j is currently at %d\n", (*b)[zeroBoard[i].x][zeroBoard[i].y], i, j)
				j--
				continue
			}
			//fmt.Println(b)
			break // Have valid value, break out of incrementing loop
		}
		//fmt.Printf("Incrementing row, i is %v\n", i)
	}
	fmt.Printf("Finished after %d steps\n", stepCount)
}

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
	//rand.Seed(1)
	rand.Seed(time.Now().UTC().UnixNano())
	p := loadCsv("./puzzles.csv")
	puzzle := rand.Intn(len(p))
	b := genBoard(p[puzzle])
	fmt.Printf("Solving puzzle number %d from file:\n\n%v\n", puzzle+2, b)
	//b := genBoard("......9.....5....85..83....35..82.....6...2....937..4.76........3.4.5.76..2......")
	//fmt.Println(p)
	start := time.Now()
	solve(&b)
	elapsed := time.Since(start)
	fmt.Printf("It took %v to run\n\n", elapsed)
	fmt.Println(b)
}

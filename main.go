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

func chkHor(x int) {

}

func chkVert(y int) {

}

func chkBox() {

}

func main() {
	fmt.Println("Hello world!")
	b := genBoard(".1....9.....5....85..83....35..82.....6...2....937..4.76........3.4.5.76..2......")
	fmt.Println(b)
}

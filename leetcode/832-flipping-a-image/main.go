package main

import (
	"fmt"
	"log"
)

func flipAndInvertImage(A [][]int) [][]int {
	for j, b := range A {
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
			if b[i] == 0 {
				b[i] = 1
			} else {
				b[i] = 0
			}

			if b[len(b)-1-i] == 0 {
				b[len(b)-1-i] = 1
			} else {
				b[len(b)-1-i] = 0
			}
		}

		if len(b)%2 != 0 {
			log.Print(len(b) / 2)
			if b[len(b)/2] == 0 {
				b[len(b)/2] = 1
			} else {
				b[len(b)/2] = 0
			}
		}
		A[j] = b
	}
	return A
}

func main() {
	var grid [][]int

	// example
	grid = append(grid, []int{1, 1, 0})
	grid = append(grid, []int{1, 0, 1})
	grid = append(grid, []int{0, 0, 0})

	fmt.Println(flipAndInvertImage(grid))
}

package main

import (
	"fmt"
)

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	n := len(matrix)
	m := 0
	for range matrix[0] {
		m++
	}
	res := make([]int, n*m)

	var x, y int
	for i := 0; i < len(res); i++ {
		res[i] = matrix[x][y]
		if (x+y)%2 == 0 {
			if y == m-1 {
				x++
			} else if x == 0 {
				y++
			} else {
				x--
				y++
			}
		} else {
			if x == n-1 {
				y++
			} else if y == 0 {
				x++
			} else {
				x++
				y--
			}
		}
	}
	return res
}

func main() {
	var grid [][]int

	// example
	grid = append(grid, []int{1, 2, 3})
	grid = append(grid, []int{4, 5, 6})

	fmt.Println(findDiagonalOrder(grid))
}

package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	for top <= bottom && left <= right {
		// traverse right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// traverse down
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// traverse left
		for i := right; i >= left; i-- {
			if top > bottom {
				continue
			}
			res = append(res, matrix[bottom][i])
		}
		bottom--
		// traverse up
		for i := bottom; i >= top; i-- {
			if left > right {
				continue
			}
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func main() {
	var grid [][]int

	// example
	grid = append(grid, []int{1, 2, 3, 4})
	grid = append(grid, []int{5, 6, 7, 8})
	grid = append(grid, []int{9, 10, 11, 12})

	fmt.Println(spiralOrder(grid))
}

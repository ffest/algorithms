package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		val :=  matrix[row][col]
		 if val > target {
			col--
		} else if val < target {
			row++
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5}, {6, 7, 8, 8, 10}, {11, 12, 13, 14, 15},
	}
	target := 4
	fmt.Println(searchMatrix(matrix, target))
}

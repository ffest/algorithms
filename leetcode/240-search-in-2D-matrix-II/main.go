package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target {
		return false
	}

	row, column := 0, len(matrix[0])-1
	for column >= 0 && row < len(matrix) {
		if matrix[row][column] > target {
			column--
		} else if matrix[row][column] < target {
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

package main

import "log"

// TODO: To be refactored
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix[0])
	var i int
	for i < len(matrix)-1 && matrix[i][m-1] < target {
		i++
	}
	for j := 0; j < m; j++ {
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15},
	}
	matrix = [][]int{}
	target := 10
	log.Print(searchMatrix(matrix, target))
}

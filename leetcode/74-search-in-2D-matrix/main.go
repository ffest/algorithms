package main

import "log"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n := len(matrix)
	m := len(matrix[0])
	if target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}

	var mid int
	low, high := 0, n-1
	for low <= high {
		mid = (low + high) / 2
		if matrix[mid][0] > target {
			high = mid - 1
		} else if matrix[mid][0] < target {
			low = mid + 1
		} else {
			return true
		}
	}

	left, right := 0, m-1
	for left <= right {
		mid = (left + right) / 2
		if matrix[high][mid] > target {
			right = mid - 1
		} else if matrix[high][mid] < target {
			left = mid + 1
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
	log.Print(searchMatrix(matrix, target))
}

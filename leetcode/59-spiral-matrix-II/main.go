package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var num = 1
	for top <= bottom && left <= right {
		// traverse right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// traverse down
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// traverse left
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// traverse up
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	fmt.Println(generateMatrix(1))
}

package main

import (
	"fmt"
	"math"
)

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	var maxSize int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

			if maxSize < dp[i][j] {
				maxSize = dp[i][j]
			}
		}
	}
	return maxSize * maxSize
}

func min(a ...int) int {
	min := math.MaxInt32
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

func main() {
	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(input))
}

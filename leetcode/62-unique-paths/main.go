package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}

func main() {
	n, m := 3, 7
	fmt.Println(uniquePaths(m, n))
}

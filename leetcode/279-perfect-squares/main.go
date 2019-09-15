package main

import (
	"fmt"
)

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	target := 12
	fmt.Println(numSquares(target))
}

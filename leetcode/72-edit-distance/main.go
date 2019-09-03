package main

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = 1 + min(dp[i][j], dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[n][m]
}

func min(a, b, c int) int {
	min := math.MaxInt32
	if a < min {
		min = a
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}

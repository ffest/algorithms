package main

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	if k >= len(prices)/2 {
		return fastSolve(prices)
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, len(prices))
	}

	for i := 1; i <= k; i++ {
		tmpMax := -prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], tmpMax+prices[j])
			tmpMax = max(tmpMax, dp[i-1][j-1]-prices[j])
		}
	}

	return dp[k][len(prices)-1]
}

func fastSolve(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(2, prices))
}

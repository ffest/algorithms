package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := -1
		for _, c := range coins {
			if i >= c && dp[i-c] != -1 {
				if min == -1 || dp[i-c]+1 < min {
					min = dp[i-c] + 1
				}
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 8
	fmt.Println(coinChange(coins, amount))
}

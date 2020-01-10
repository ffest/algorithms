package main

import (
	"fmt"
)

// O(n) additional space
/*func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}
	return min(dp[len(cost)-2], dp[len(cost)-1])
}
*/
// No additional space required
func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-2], cost[i-1])
	}
	return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost := []int{0, 1, 1, 0}
	fmt.Println(minCostClimbingStairs(cost))
}

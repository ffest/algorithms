package main

import (
	"fmt"
)

var mod = 1000000007

// Slow backtracking solution
/*func numRollsToTarget(d int, f int, target int) int {
	if d*f < target {
		return 0
	}
	var count int
	backtracking(d, f, target, &count)
	return count
}

func backtracking(d, f, target int, count *int) {
	if target < 0 || d < 0 {
		return
	}
	if d == 0 && target == 0 {
		*count++
		*count = *count % mod
		return
	}
	for j := 1; j <= f; j++ {
		backtracking(d-1, f, target-j, count)
	}
}*/

// Fast DP solution
func numRollsToTarget(d int, f int, target int) int {
	if d*f < target {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= d; i++ {
		dp1 := make([]int, target+1)
		for j := 1; j <= f; j++ {
			for k := j; k <= target; k++ {
				dp1[k] = (dp1[k] + dp[k-j]) % mod
			}
		}
		dp = dp1
	}
	return dp[target]
}

func main() {
	d := 30
	f := 30
	target := 500
	fmt.Println(numRollsToTarget(d, f, target))
}

package main

import (
	"fmt"
)

func minTaps(n int, ranges []int) int {
	dp := make([]int, n+1)
	for i := range ranges {
		if ranges[i] == 0 {
			continue
		}
		interval := []int{i - ranges[i], i + ranges[i]}
		if interval[0] < 0 {
			interval[0] = 0
		}
		if interval[1] > n {
			interval[1] = n
		}
		taps := dp[interval[0]]
		for j := interval[0]; j <= interval[1]; j++ {
			if dp[j] == 0 {
				if interval[0] == 0 {
					dp[j] = 1
				} else {
					dp[j] = taps + 1
				}
			} else {
				dp[j] = min(dp[j], taps+1)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 3
	ranges := []int{0, 0, 0, 0}
	fmt.Println(minTaps(n, ranges))
}

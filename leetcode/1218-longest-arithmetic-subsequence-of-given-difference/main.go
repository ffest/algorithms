package main

import (
	"fmt"
)

func longestSubsequence(arr []int, difference int) int {
	if len(arr) == 0 {
		return 0
	}
	var max int
	cache := make(map[int]int)
	for i := range arr {
		cache[arr[i]] = 1 + cache[arr[i]-difference]
		if cache[arr[i]] > max {
			max = cache[arr[i]]
		}
	}
	return max
}

// O(n^2) solution. Too slow
/*func longestSubsequence(arr []int, difference int) int {
	if len(arr) == 0 {
		return 0
	}

	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}
	max := 1
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] != difference {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return max
}*/

func main() {
	arr := []int{3, 4, -3, -2, -4}
	diff := -5
	fmt.Println(longestSubsequence(arr, diff))
}

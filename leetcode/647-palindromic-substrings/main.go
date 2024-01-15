package main

import (
	"fmt"
)

// DP solution based on Leetcode 5
/*func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var counter int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if (j-i < 2 || dp[i+1][j-1]) && s[i] == s[j] {
				dp[i][j] = true
				counter++
			}
		}
	}
	return counter
}*/

// Expand substr
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	var counter int
	for i := 0; i < len(s); i++ {
		expand(i, i, s, &counter)
		expand(i, i+1, s, &counter)
	}
	return counter
}

func expand(left, right int, s string, counter *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*counter++
		left--
		right++
	}
}

func main() {
	s := "abcac"
	fmt.Println(countSubstrings(s))
}

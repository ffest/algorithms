package main

import (
	"fmt"
)

// DP solution
/*func longestPalindrome(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	}

	dp := make([][]bool, len(str))
	for i := range dp {
		dp[i] = make([]bool, len(str))
	}

	var maxPalindpom string
	var maxLength int

	// case for length of 1
	for i := range str {
		dp[i][i] = true
	}
	maxPalindpom = str[:1]
	maxLength = 1

	// case for length of 2
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			maxPalindpom = str[i-1 : i+1]
			maxLength = 2
			dp[i-1][i] = true
		}
	}
	// case for length of greater than 2
	for i := 2; i < len(str); i++ {
		for j := 0; j <= i-2; j++ {
			if str[i] == str[j] && dp[j+1][i-1] {
				if maxLength < i-j+1 {
					maxPalindpom = str[j : i+1]
					maxLength = i - j + 1
				}

				dp[j][i] = true
			}
		}
	}

	return maxPalindpom
}*/

// Expand solution
func longestPalindrome(str string) string {
	var max string
	for i := 0; i < len(str); i++ {
		expand(i, i, str, &max)
		expand(i, i+1, str, &max)
	}
	return max
}

func expand(left, right int, s string, max *string) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		if len(*max) < right-left+1 {
			*max = s[left : right+1]
		}
		left--
		right++
	}
}

func main() {
	fmt.Println(longestPalindrome("cababac"))
}

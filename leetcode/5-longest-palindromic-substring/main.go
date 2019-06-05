package main

import (
	"fmt"
)

func longestPalindrome(str string) string {
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
	for i := 0; i < len(str); i++ {
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
				dp[j][i] = true
				if i-j+1 > maxLength {
					maxLength = i - j + 1
					maxPalindpom = str[j : i+1]
				}
			}
		}
	}

	return maxPalindpom
}

func main() {
	fmt.Println(longestPalindrome("cababac"))
}

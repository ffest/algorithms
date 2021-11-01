package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if !dp[i] {
			continue
		}
		for _, word := range wordDict {
			if len(word) <= len(s)-i && word == s[i:i+len(word)] {
				dp[i+len(word)] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "leetcode"
	dict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, dict))
}

package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		if !dp[i] {
			continue
		}

		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i+len(w)] = true
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

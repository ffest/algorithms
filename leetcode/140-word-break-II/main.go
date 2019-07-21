package main

import (
	"fmt"
)

// Too slow
/*func wordBreak(s string, wordDict []string) []string {
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	var tmp string
	for i := range s {
		log.Print(i)
		if dp[i] == nil {
			continue
		}
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				for _, str := range dp[i] {
					if len(str) == 0 {
						tmp = w
					} else {
						tmp = str + " " + w
					}
					dp[i+len(w)] = append(dp[i+len(w)], tmp)
				}
			}
		}
	}
	return dp[len(s)]
}*/

func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, map[string][]string{})
}

func dfs(s string, wordDict []string, cache map[string][]string) []string {
	if res, ok := cache[s]; ok {
		return res
	}
	if len(s) == 0 {
		return []string{""}
	}

	var res []string
	for _, w := range wordDict {
		if len(w) <= len(s) && w == s[:len(w)] {
			for _, str := range dfs(s[len(w):], wordDict, cache) {
				if len(str) == 0 {
					res = append(res, w)
				} else {
					res = append(res, w+" "+str)
				}
			}
		}
	}
	cache[s] = res
	return res
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	words := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, words))
}

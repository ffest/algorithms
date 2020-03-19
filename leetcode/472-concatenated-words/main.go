package main

import (
	"fmt"
	"sort"
)

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	result := make([]string, 0)
	usedWords := make(map[string]struct{})
	for i := 0; i < len(words); i++ {
		if canForm(words[i], usedWords) {
			result = append(result, words[i])
		}
		usedWords[words[i]] = struct{}{}
	}
	return result
}

func canForm(s string, words map[string]struct{}) bool {
	if len(words) == 0 {
		return false
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if !dp[j] {
				continue
			}
			if _, ok := words[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}

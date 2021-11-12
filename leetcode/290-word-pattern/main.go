package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	patternMap := [26]string{}
	wordUsed := make(map[string]struct{})
	for i, w := range words {
		match := patternMap[pattern[i]-'a']
		if match == "" {
			if _, ok := wordUsed[w]; ok {
				return false
			}
			wordUsed[w] = struct{}{}
			patternMap[pattern[i]-'a'] = w
			continue
		} else if match != w {
			return false
		}
	}
	return true
}

func main() {
	pattern, s := "abba", "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}

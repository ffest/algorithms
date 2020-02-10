package main

import (
	"fmt"
)

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var max, i, j int
	chars := make(map[uint8]int)
	for i = 0; i < len(s); i++ {
		chars[s[i]]++
		for len(chars) > 2 {
			chars[s[j]]--
			if chars[s[j]] == 0 {
				delete(chars, s[j])
			}
			j++
		}
		if i-j+1 > max {
			max = i - j + 1
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	cache := [128]bool{}
	var max, length, tail int
	for i := range s {
		for cache[s[i]] {
			cache[s[tail]] = false
			length--
			tail++
		}
		cache[s[i]] = true
		length++
		if max < length {
			max = length
		}
	}
	return max
}

func main() {
	input := "bbb"
	fmt.Println(lengthOfLongestSubstring(input))
}

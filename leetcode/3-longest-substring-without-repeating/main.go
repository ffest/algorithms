package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var m [128]bool
	var max, tail, length int
	for i := range s {
		for m[s[i]] {
			m[s[tail]] = false
			tail++
			length--
		}
		m[s[i]] = true
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

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var m [128]bool
	var max int

	l := 0
	var current int
	for i := range s {
		char := s[i]
		for m[char] {
			m[s[current]] = false
			current++
			l--
		}
		m[char] = true
		l++
		if max < l {
			max = l
		}
	}
	return max
}

func main() {
	input := "bbb"
	fmt.Println(lengthOfLongestSubstring(input))
}

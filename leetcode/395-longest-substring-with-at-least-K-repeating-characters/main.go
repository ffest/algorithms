package main

import (
	"fmt"
)

func longestSubstring(s string, k int) int {
	var maxLength int
	var idx uint8

	for h := 1; h <= 26; h++ {
		cache := [26]int{}
		i, j, noLessThanK, unique := 0, 0, 0, 0
		for j < len(s) {
			if unique <= h {
				idx = s[j] - 'a'
				if cache[idx] == 0 {
					unique++
				}
				cache[idx]++
				if cache[idx] == k {
					noLessThanK++
				}
				j++
			} else {
				idx = s[i] - 'a'
				if cache[idx] == k {
					noLessThanK--
				}
				cache[idx]--
				if cache[idx] == 0 {
					unique--
				}
				i++
			}
			if unique == h && unique == noLessThanK && j-i > maxLength {
				maxLength = j - i
			}
		}
	}

	return maxLength
}

func main() {
	fmt.Println(longestSubstring("bbaaacbd", 3))
}

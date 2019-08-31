package main

import (
	"fmt"
	"math"
)

func firstUniqChar(s string) int {
	cache := make(map[int32]int)
	for i, l := range s {
		if _, ok := cache[l]; ok {
			cache[l] = -1
			continue
		}
		cache[l] = i
	}

	firstLetter := math.MaxInt32
	for _, v := range cache {
		if v < firstLetter && v >= 0 {
			firstLetter = v
		}
	}
	if firstLetter == math.MaxInt32 {
		return -1
	}
	return firstLetter
}

func main() {
	input := "leetcode"
	fmt.Println(firstUniqChar(input))
}

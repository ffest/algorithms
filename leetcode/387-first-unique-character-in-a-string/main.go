package main

import (
	"fmt"
)

// solution based on map
/*func firstUniqChar(s string) int {
	cache := make(map[int32]int)
	for i, l := range s {
		if _, ok := cache[l]; ok {
			cache[l] = -1
			continue
		}
		cache[l] = i
	}

	first := math.MaxInt32
	for _, v := range cache {
		if v != -1 && first > v {
			first = v
		}
	}
	if first == math.MaxInt32 {
		return -1
	}
	return first
}*/

// solution based on [26]int{}
func firstUniqChar(s string) int {
	cache := [26]int{}
	for _, l := range s {
		cache[l-'a']++
	}

	for i := range s {
		if cache[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	input := "leetcode"
	fmt.Println(firstUniqChar(input))
}

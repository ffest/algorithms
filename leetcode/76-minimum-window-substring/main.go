package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	var result string
	cache := [128]int{}
	for i := range t {
		cache[t[i]]++
	}

	var left, right int
	counter := len(t)
	for right = 0; right < len(s); right++ {
		if cache[s[right]] > 0 {
			counter--
		}
		cache[s[right]]--
		for counter == 0 {
			if result == "" || right-left < len(result) {
				result = s[left : right+1]
			}
			cache[s[left]]++
			if cache[s[left]] > 0 {
				counter++
			}
			left++
		}
	}
	return result
}

func main() {
	s := "abac"
	t := "bc"
	fmt.Println(minWindow(s, t))
}

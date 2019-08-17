package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	m := [128]int{}
	for i := range t {
		m[t[i]]++
	}

	var result string
	var left, right int
	counter := len(t)
	for right = 0; right < len(s); right++ {
		if m[s[right]] > 0 {
			counter--
		}
		m[s[right]]--

		for counter == 0 {
			if result == "" || right-left < len(result) {
				result = s[left : right+1]
			}
			m[s[left]]++
			if m[s[left]] > 0 {
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

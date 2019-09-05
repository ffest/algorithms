package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	result := make([]int, 0)
	need, have := [26]int{}, [26]int{}
	for i := range p {
		need[p[i]-'a']++
		have[s[i]-'a']++
	}
	var tail int
	for i := len(p); i < len(s); i++ {
		if have == need {
			result = append(result, tail)
		}
		have[s[tail]-'a']--
		have[s[i]-'a']++
		tail++
	}
	if have == need {
		result = append(result, tail)
	}

	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

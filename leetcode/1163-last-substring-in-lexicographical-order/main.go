package main

import (
	"fmt"
	"sort"
)

func lastSubstring(s string) string {
	var maxChar = 'a'
	for _, c := range s {
		if c > maxChar {
			maxChar = c
		}
	}

	positions := make([]int, 0)
	for i, c := range s {
		if c == maxChar {
			positions = append(positions, i)
		}
	}

	strs := make([]string, 0, len(positions))
	for _, p := range positions {
		strs = append(strs, s[p:])
	}
	sort.Strings(strs)
	return strs[len(strs)-1]
}

func main() {
	s := "abab"
	fmt.Println(lastSubstring(s))
}

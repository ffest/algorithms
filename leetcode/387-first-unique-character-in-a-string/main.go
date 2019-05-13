package main

import "fmt"

func firstUniqChar(s string) int {
	strChars := make(map[int32]int)
	for i, c := range s {
		if _, ok := strChars[c]; ok {
			strChars[c] = -1
			continue
		}
		strChars[c] = i
	}
	var min = -1
	for _, c := range strChars {
		if c == -1 {
			continue
		}
		if min == -1 || min > c {
			min = c
		}
	}
	return min
}

func main() {
	input := "leetcode"
	fmt.Println(firstUniqChar(input))
}

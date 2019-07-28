package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var start int
	for end := range s {
		if s[end] == ' ' {
			s = s[:start] + reverseStr(s[start:end]) + s[end:]
			end++
			start = end
		}
	}
	s = s[:start] + reverseStr(s[start:])
	return s
}

func reverseStr(s string) string {
	str := []byte(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}
	return string(str)
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

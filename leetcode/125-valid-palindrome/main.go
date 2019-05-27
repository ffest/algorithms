package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start, end := 0, len(s)-1
	for start <= end {
		if (s[start] < 'a' || s[start] > 'z') && (s[start] < '0' || s[start] > '9') {
			start++
			continue
		}
		if (s[end] < 'a' || s[end] > 'z') && (s[end] < '0' || s[end] > '9') {
			end--
			continue
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	input := "0P"
	fmt.Println(isPalindrome(input))
}

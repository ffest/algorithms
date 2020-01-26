package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "bbaabaaa"
	fmt.Println(removePalindromeSub(s))
}

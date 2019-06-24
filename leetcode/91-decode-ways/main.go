package main

import (
	"fmt"
)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	count := 1
	prevCount := 1
	for i := 1; i < len(s); i++ {
		currCount := 0
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			currCount += prevCount
		}

		if s[i] > '0' {
			currCount += count
		}

		prevCount = count
		count = currCount
	}

	return count
}

func main() {
	input := "1234321"
	fmt.Println(numDecodings(input))
}

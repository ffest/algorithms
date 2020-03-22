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

	prev, count := 1, 1
	for i := 1; i < len(s); i++ {
		current := 0
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			current += prev
		}
		if s[i] > '0' {
			current += count
		}
		prev = count
		count = current
	}
	return count
}

func main() {
	input := "1234321"
	fmt.Println(numDecodings(input))
}

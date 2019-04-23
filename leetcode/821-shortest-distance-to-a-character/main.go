package main

import (
	"fmt"
)

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))

	var prev = -1
	for i := range s {
		if s[i] == c {
			for j := prev + 1; j < i; j++ {
				if res[j] == 0 || res[j] > i-j {
					res[j] = i - j
				}
			}
			prev = i
			res[i] = 0
		} else if prev != -1 {
			res[i] = i - prev
		}

	}

	return res
}

func main() {
	input := "loveleetcode"
	char := byte('e')
	fmt.Println(shortestToChar(input, char))
}

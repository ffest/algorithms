package main

import (
	"fmt"
)

func maxProduct(words []string) int {
	var max int
	masks := make([]int64, len(words))
	for i := range words {
		for j := range words[i] {
			masks[i] |= 1 << (words[i][j] - 'a')
		}
	}
	for i := range masks {
		for j := range masks {
			if masks[i]&masks[j] == 0 && max < len(words[i])*len(words[j]) {
				max = len(words[i]) * len(words[j])
			}
		}
	}
	return max
}

func main() {
	input := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(input))
}

package main

import (
	"fmt"
)

func maxProduct(words []string) int {
	var max int
	masks := make([]int, len(words))
	for i, word := range words {
		for j := range word {
			masks[i] |= 1 << (word[j] - 'a')
		}
	}

	for i := 0; i < len(masks); i++ {
		for j := 0; j < len(masks); j++ {
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

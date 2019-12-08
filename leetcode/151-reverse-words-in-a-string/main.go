package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var idx int
	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			words[idx] = words[i]
			idx++
		}
	}
	words = words[:idx]
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example  "))
}

package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	words := strings.Split(s, " ")
	maxLength := 0
	for _, w := range words {
		if maxLength < len(w) {
			maxLength = len(w)
		}
	}
	result := make([]string, maxLength)
	for i := 0; i < maxLength; i++ {
		element := ""
		for j := range words {
			if words[j] == "" {
				element += " "
			} else {
				element += string(words[j][0])
				words[j] = words[j][1:]
			}
		}
		result[i] = strings.TrimRight(element, " ")
	}
	return result
}

func main() {
	s := "CONTEST IS COMING"
	fmt.Println(printVertically(s))
}

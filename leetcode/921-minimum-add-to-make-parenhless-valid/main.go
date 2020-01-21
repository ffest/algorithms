package main

import (
	"fmt"
)

func minAddToMakeValid(S string) int {
	var lefts, rights int
	for i := range S {
		if S[i] == '(' {
			lefts++
		} else {
			if lefts > 0 {
				lefts--
			} else {
				rights++
			}
		}
	}
	return lefts + rights
}

func main() {
	input := "(()())(("
	fmt.Println(minAddToMakeValid(input))
}

package main

import (
	"fmt"
)

func minAddToMakeValid(S string) int {
	var lefts int
	var rights int
	for _, r := range S {
		if r == '(' {
			lefts++
		} else {
			if lefts > 0 {
				lefts -= 1
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

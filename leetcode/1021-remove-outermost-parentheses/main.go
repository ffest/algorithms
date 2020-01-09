package main

import "fmt"

func removeOuterParentheses(S string) string {
	var result string
	var brackets, prev int
	for i := range S {
		if S[i] == '(' {
			brackets++
		} else {
			brackets--
		}
		if brackets == 0 {
			result += S[prev+1 : i]
			prev = i + 1
		}
	}
	return result
}

func main() {
	s := "()()"
	fmt.Println(removeOuterParentheses(s))
}

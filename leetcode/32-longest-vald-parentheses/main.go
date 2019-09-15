package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	var result int
	stack := make([]int, 0)
	for i := range s {
		if len(stack) > 0 && s[i] == ')' && s[stack[len(stack)-1]] == '(' {
			stack = stack[:len(stack)-1]
			if result < i-stack[len(stack)-1] {
				result = i - stack[len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}
	return result
}

func main() {
	//input := "(()()))()()"
	input := ")()())"
	fmt.Println(longestValidParentheses(input))
}

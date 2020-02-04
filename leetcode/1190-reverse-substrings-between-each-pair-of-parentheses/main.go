package main

import "fmt"

func reverseParentheses(s string) string {
	current, stack := make([]byte, 0), make([][]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, current)
			current = make([]byte, 0)
		} else if s[i] == ')' {
			reverse(current)
			if len(stack) > 0 {
				current = append(stack[len(stack)-1], current...)
				stack = stack[:len(stack)-1]
			}
		} else {
			current = append(current, s[i])
		}
	}
	return string(current)
}

func reverse(s []byte) {
	for i := 0; i < (len(s))/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	s := "a(bcdefghijkl(mno)p)q"
	fmt.Println(reverseParentheses(s))
}

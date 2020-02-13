package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	result := make([]byte, 0)
	opens := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opens++
		} else if s[i] == ')' {
			if opens == 0 {
				continue
			}
			opens--
		}
		result = append(result, s[i])
	}

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == '(' && opens > 0 {
			opens--
			result = append(result[:i], result[i+1:]...)
		}
	}

	return string(result)
}

func main() {
	s := "(a(b(c)d)"
	fmt.Println(minRemoveToMakeValid(s))
}

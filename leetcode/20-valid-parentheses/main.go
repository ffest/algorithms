package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s)%2 == 1 {
		return false
	}

	closes := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := range s {
		if opens, ok := closes[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opens {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	input := "()[]{}"
	fmt.Println(isValid(input))
}

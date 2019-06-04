package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	closedBrackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)
	for i := range s {
		if openingBracket, ok := closedBrackets[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openingBracket {
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

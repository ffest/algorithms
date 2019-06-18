package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	var num1, num2 int
	for _, token := range tokens {
		if len(token) == 1 && strings.ContainsAny(token, "+-*/") {
			num2 = stack[len(stack)-1]
			num1 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	input := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(input))
}

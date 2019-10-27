package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]byte, len(num1)+len(num2))

	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			tmp := (num2[i]-'0')*(num1[j]-'0') + result[i+j+1]
			result[i+j+1] = tmp % 10
			result[i+j] += tmp / 10
		}
	}

	var zeros int
	for result[zeros] == 0 {
		zeros++
	}
	for i := zeros; i < len(result); i++ {
		result[i] += '0'
	}
	return string(result[zeros:])
}

func main() {
	num1 := "234"
	num2 := "325"
	fmt.Println(multiply(num1, num2))
}

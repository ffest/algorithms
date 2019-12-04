package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	} else if num2 == "0" {
		return num1
	}

	var carry uint8
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	result := make([]byte, len(num1))
	i, j := len(num1)-1, len(num2)-1
	for j >= 0 {
		tmp := num1[i] + num2[j] - 2*'0' + carry
		result[i] = tmp % 10
		carry = tmp / 10
		i--
		j--
	}
	for i >= 0 {
		tmp := num1[i] - '0' + carry
		result[i] = tmp % 10
		carry = tmp / 10
		i--
	}
	for i := 0; i < len(result); i++ {
		result[i] += '0'
	}
	if carry > 0 {
		return "1" + string(result)
	}
	return string(result)
}

func main() {
	num1 := "1"
	num2 := "999"
	fmt.Println(addStrings(num1, num2))
}

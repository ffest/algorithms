package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	var result, i int
	var sign = 1
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign *= -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && s[i] > '7') {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + int(s[i]-'0')
		i++
	}

	return result * sign
}

func myItoa(n int) string {
	var result string
	var sign = 1
	if n < 0 {
		sign *= -1
		n *= -1
	}
	for n > 0 {
		result = string(rune(n%10+'0')) + result
		n /= 10
	}

	if sign < 0 {
		return "-" + result
	}
	return result
}

func main() {
	fmt.Println(myAtoi("   -42"))
	fmt.Println(-42)
}

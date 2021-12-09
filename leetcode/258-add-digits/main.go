package main

import "fmt"

func addDigits(num int) int {
	var result int
	for num > 0 {
		result += num % 10
		num /= 10
		if num == 0 && result >= 10 {
			num = result
			result = 0
		}
	}
	return result
}

func main() {
	fmt.Println(addDigits(38))
}

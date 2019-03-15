package main

import "fmt"

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for {
		if i < 0 {
			digits := append([]int{1}, digits...)
			return digits
		}
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
		i--
	}

	return digits
}

func main() {
	input := []int{9, 9, 9, 9}
	fmt.Println(plusOne(input))
}

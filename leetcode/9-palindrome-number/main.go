package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp := x
	var reversed int
	for tmp > 0 {
		reversed = reversed*10 + tmp%10
		tmp /= 10
	}

	return reversed == x
}

func main() {
	x := 1
	fmt.Println(isPalindrome(x))
}

package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	// g - guess
	g := num
	for g*g > num {
		g = (g + num/g) / 2
	}
	return g*g == num
}

func main() {
	input := 15
	fmt.Println(isPerfectSquare(input))
}

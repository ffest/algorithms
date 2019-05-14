package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	one, two := 1, 1
	for i := 2; i < n; i++ {
		next := one + two
		one = two
		two = next
	}
	return one + two
}

func main() {
	input := 6
	fmt.Println(climbStairs(input))
}

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		x = -x
		negative = true
	}

	var result int
	for x > 0 {
		result = result*10 + x%10
		if result > math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	if negative {
		result = -result
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
}

package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var sign int
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	} else {
		sign = -1
	}

	result, power := uint(0), uint(32)
	dividend, divisor = abs(dividend), abs(divisor)
	divPower := uint(divisor) << power

	for dividend >= divisor {
		for divPower > uint(dividend) {
			divPower >>= 1
			power--
		}
		result += 1 << power
		dividend -= int(divPower)
	}

	return int(result) * sign
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}

package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}

	length := int(math.Log10(float64(x)))
	mask := int(math.Pow10(length))
	for i := 0; i <= length/2; i++ {
		if x%10 != x/mask {
			return false
		}
		x %= mask
		x /= 10
		mask /= 100
	}

	return true
}

/*func isPalindrome(x int) bool {
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
}*/

func main() {
	x := 12321
	fmt.Println(isPalindrome(x))
}

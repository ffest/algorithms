package main

import "fmt"

func isHappy1(n int) bool {
	a := helper(n)
	b := helper(a)
	for a != 1 && a != b {
		a = helper(a)
		b = helper(helper(b))
	}
	return a == 1
}

func isHappy(n int) bool {
	for {
		if n == 1 {
			return true
		} else if n == 4 {
			return false
		}
		n = helper(n)
	}
}

func helper(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy1(19))
	fmt.Println(isHappy(19))
}

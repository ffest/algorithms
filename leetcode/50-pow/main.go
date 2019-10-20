package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x, n = x*x, n>>1
	}
	return result
}

func main() {
	fmt.Println(myPow(2.0, -2))
}

package main

import (
	"fmt"
)

func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		if !zeroNum(i) && !zeroNum(n-i) {
			return []int{i, n - i}
		}
	}
	return nil
}

func zeroNum(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}

func main() {
	n := 2
	fmt.Println(getNoZeroIntegers(n))
}

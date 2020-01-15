package main

import (
	"fmt"
)

func sumSubarrayMins(A []int) int {
	var count, mod int
	mod = 1e9 + 7
	lefts, rights := make([]int, len(A)), make([]int, len(A))

	stack := make([]int, 0)
	for i := 0; i < len(A); i++ {
		for len(stack) > 0 && A[stack[len(stack)-1]] > A[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lefts[i] = i + 1
		} else {
			lefts[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := len(A) - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack)-1]] >= A[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rights[i] = len(A) - i
		} else {
			rights[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(A); i++ {
		count += A[i] * lefts[i] * rights[i]
		count %= mod
	}

	return count
}

func main() {
	A := []int{71, 55, 82, 55}
	fmt.Println(sumSubarrayMins(A))
}

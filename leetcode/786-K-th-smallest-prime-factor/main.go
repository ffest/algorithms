package main

import (
	"fmt"
)

func kthSmallestPrimeFraction(A []int, K int) []int {
	left, right := 0.0, 1.0
	p, q := 0, 1

	n := len(A)
	for {
		p = 0
		count := 0

		mid := (left + right) / 2

		for i, j := 0, n-1; i < n; i++ {
			for j >= 0 && float64(A[i]) > mid*float64(A[n-1-j]) {
				j--
			}
			count += j + 1
			if j >= 0 && p*A[n-1-j] < q*A[i] {
				p, q = A[i], A[n-1-j]
			}
		}

		if count < K {
			left = mid
		} else if count > K {
			right = mid
		} else {
			return []int{p, q}
		}
	}
}

func main() {
	A := []int{1, 2, 3, 5}
	fmt.Println(kthSmallestPrimeFraction(A, 3))
}

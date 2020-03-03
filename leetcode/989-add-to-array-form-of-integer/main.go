package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	var carry int
	idx := len(A) - 1
	for idx >= 0 || K > 0 {
		tmp := carry
		if K > 0 {
			tmp += K % 10
		}
		if idx >= 0 {
			tmp += A[idx]
			A[idx] = tmp % 10
			idx--
		} else {
			A = append([]int{tmp % 10}, A...)
		}
		carry = tmp / 10
		K /= 10
	}
	if carry == 1 {
		A = append([]int{1}, A...)
	}
	return A
}

func main() {
	A := []int{2, 1, 5}
	K := 1806
	fmt.Println(addToArrayForm(A, K))
}

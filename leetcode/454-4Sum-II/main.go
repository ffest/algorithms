package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	cache := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			cache[A[i]+B[j]]++
		}
	}

	var count int
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			target := -(C[i] + D[j])
			if c, ok := cache[target]; ok {
				count += c
			}
		}
	}
	return count
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}

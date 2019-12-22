package main

import (
	"fmt"
)

func pancakeSort(A []int) []int {
	flips := make([]int, 0)
	sortedIndex := len(A)
	var max, maxIndex int
	for sortedIndex > 1 {
		for i := 0; i < sortedIndex; i++ {
			if A[i] > max {
				max = A[i]
				maxIndex = i
			}
		}
		if maxIndex != sortedIndex-1 {
			flips = append(flips, maxIndex+1)
			flip(A, maxIndex+1)
			flips = append(flips, sortedIndex)
			flip(A, sortedIndex)
		}
		sortedIndex--
		max = 0
	}
	return flips
}

func flip(A []int, k int) {
	left, right := 0, k-1
	for left < right {
		A[left], A[right] = A[right], A[left]
		left++
		right--
	}
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(pancakeSort(A))
}

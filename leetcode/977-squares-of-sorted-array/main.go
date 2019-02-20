package main

import (
	"fmt"
	"math/rand"
)

func sortedSquares(A []int) []int {
	for index, element := range A {
		A[index] = element * element
	}

	return Quick(A)
}

// Quicksort  is an efficient sorting algorithm, serving as a systematic method
// for placing the elements of an array in order
// On average, the algorithm takes O(n log n) comparisons to sort n items.
// In the worst case, it makes O(n2) comparisons, though this behavior is rare.
func Quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := 0, len(arr)-1
	median := rand.Int() % len(arr)

	arr[median], arr[right] = arr[right], arr[median]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Quick(arr[:left])
	Quick(arr[left+1:])

	return arr
}

func main() {
	A := []int{-4, -2, 1, 3, 11}
	fmt.Println(sortedSquares(A))
}

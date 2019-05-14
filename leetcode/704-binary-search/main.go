package main

import "fmt"

func search(arr []int, n int) int {
	left := 0
	right := len(arr)
	for {
		median := (left + right) / 2
		switch {
		case left == right:
			return -1
		case n < arr[median]:
			right = median
			continue
		case n > arr[median]:
			left = median + 1
			continue
		default:
			return median
		}
	}
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

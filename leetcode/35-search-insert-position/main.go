package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for {
		median := (left + right) / 2
		switch {
		case left == right:
			return median
		case target < nums[median]:
			right = median
			continue
		case target > nums[median]:
			left = median + 1
			continue
		default:
			return median
		}
	}
}

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(searchInsert(nums, target))
}

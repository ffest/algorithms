package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		median := (left + right) / 2
		if nums[median] == target {
			return median
		} else if target < nums[median] {
			right = median - 1
		} else {
			left = median + 1
		}
	}
	return left
}

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(searchInsert(nums, target))
}

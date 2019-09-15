package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	left, right, start, end := 0, len(nums)-1, 0, 0
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	start = left
	left, right = 0, len(nums)-1
	for left < right {
		mid := (left+right)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	end = right

	return []int{start, end}
}

func main() {
	nums := []int{2, 3, 4, 4, 4, 8}
	target := 4
	fmt.Println(searchRange(nums, target))
}

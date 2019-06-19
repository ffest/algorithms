package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	var left, right, middle, start, end int
	left, right = 0, len(nums)-1
	for left < right {
		middle = (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	start = left
	left, right = 0, len(nums)-1
	for left < right {
		middle = (left+right)/2 + 1
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle
		}
	}
	end = right
	return []int{start, end}
}

func main() {
	nums := []int{2, 3, 4, 4, 8}
	target := 4
	fmt.Println(searchRange(nums, target))
}

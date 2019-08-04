package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	//input := []int{4, 5, 6, 7, 0, 1, 2}
	//input := []int{14, 0, 2, 3, 5, 9, 12}
	input := []int{5, 1, 3}
	//input := []int{1}
	target := 0
	fmt.Println(search(input, target))
}

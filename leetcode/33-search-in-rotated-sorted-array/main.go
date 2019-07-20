package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	idx := searchShift(nums)
	var res int
	if idx == 0 {
		res = binary(nums, target)
	} else if nums[0] <= target {
		res = binary(nums[:idx], target)
	} else {
		res = binary(nums[idx:], target)
		if res != -1 {
			res += idx
		}
	}
	return res
}

func searchShift(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
	}
	return left
}

func binary(arr []int, target int) int {
	left := 0
	right := len(arr)
	for {
		median := (left + right) / 2
		switch {
		case left == right:
			return -1
		case target < arr[median]:
			right = median
			continue
		case target > arr[median]:
			left = median + 1
			continue
		default:
			return median
		}
	}
}

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	//input := []int{14, 0, 2, 3, 5, 9, 12}
	//input := []int{1}
	target := 0
	fmt.Println(search(input, target))
}

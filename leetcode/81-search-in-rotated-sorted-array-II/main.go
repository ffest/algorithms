package main

import "fmt"

func search(nums []int, target int) bool {
	mid, left, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[left] < nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			left++
		}
	}
	return false
}

func main() {
	nums := []int{1, 3, 1, 1, 1}
	target := 3
	fmt.Println(search(nums, target))
}

package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	result, lastMax, lastMin := nums[0], 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			lastMax, lastMin = lastMin, lastMax
		}
		lastMax, lastMin = max(lastMax*nums[i], nums[i]), min(lastMin*nums[i], nums[i])
		if lastMax > result {
			result = lastMax
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

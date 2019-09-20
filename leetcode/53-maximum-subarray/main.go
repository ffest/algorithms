package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(maxSubArray(input))
}

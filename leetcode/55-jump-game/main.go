package main

import (
	"fmt"
)

// refactor it
func canJump(nums []int) bool {
	var maxJump int
	for i := range nums {
		if maxJump < i {
			return false
		}
		maxJump = max(maxJump, i+nums[i])
	}
	return maxJump >= len(nums)-1
}

func canJumpReverse(nums []int) bool {
	currentStep := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if currentStep <= i+nums[i] {
			currentStep = i
		}
	}
	return currentStep == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := []int{3, 0, 8, 2, 0, 0, 1}
	fmt.Println(canJumpReverse(input))
}

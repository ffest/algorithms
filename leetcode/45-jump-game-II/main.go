package main

import (
	"fmt"
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var maxNextStep, localMaxIdx, jumps int
	var localMax = -(1<<31 - 1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxNextStep = nums[i]
			jumps++
			continue
		}

		if nums[i] <= maxNextStep-i && i != maxNextStep {
			continue
		}

		if localMax < (nums[i] - maxNextStep + i) {
			localMax = nums[i] - maxNextStep + i
			localMaxIdx = i
		}
		if i == maxNextStep && i != len(nums)-1 {
			jumps++
			maxNextStep = nums[localMaxIdx] + localMaxIdx
			localMax = -(1<<31 - 1)
		}
	}
	return jumps
}

func main() {
	//input := []int{3, 2, 1, 3, 4, 5, 1, 1, 5, 3, 1, 5}
	//input := []int{2, 3, 1, 1, 4}
	input := []int{1, 0}
	fmt.Println(jump(input))
}

package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var count = 1
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != current {
			nums[count] = nums[i]
			count++
			current = nums[i]
		}
	}
	return count
}

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(input))
}

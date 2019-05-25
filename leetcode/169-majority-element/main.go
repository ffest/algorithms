package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	result := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			result = nums[i]
		} else if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	return result
}

func main() {
	input := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(input))
}

package main

import (
	"fmt"
	"math"
)

// O(n^2) solution
/*func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for i := 0; i < len(nums)-2; i++ {
		min, max := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < max && nums[j] > min {
				return true
			} else if nums[j] > max {
				max = nums[j]
			}
		}
	}

	return false
}*/

// Solution using stack
func find132pattern(nums []int) bool {
	max := math.MinInt32
	stack := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			max = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	nums := []int{3, 5, 0, 3, 4}
	fmt.Println(find132pattern(nums))
}

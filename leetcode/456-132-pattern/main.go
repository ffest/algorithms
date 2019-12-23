package main

import (
	"fmt"
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
	if len(nums) < 3 {
		return false
	}
	return false
}

func main() {
	nums := []int{3, 5, 0, 3, 4}
	fmt.Println(find132pattern(nums))
}

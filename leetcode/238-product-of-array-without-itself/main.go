package main

import (
	"fmt"
)

/*func productExceptSelfWithDivision(nums []int) []int {
	result := make([]int, len(nums))
	zeroCounts, zeroIdx := 0, -1
	sum := 1
	for i := range nums {
		if nums[i] == 0 {
			if zeroCounts == 1 {
				return result
			}

			zeroCounts++
			zeroIdx = i
			continue
		}
		sum *= nums[i]
	}
	if zeroCounts > 0 {
		result[zeroIdx] = sum
		return result
	}

	for i := range nums {
		if i == zeroIdx {
			continue
		}
		result[i] = sum / nums[i]
	}
	return result
}*/

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	sum := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= sum
		sum *= nums[i]
	}

	return result
}

func main() {
	input := []int{2, 3, 4, 5}
	fmt.Println(productExceptSelf(input))
}

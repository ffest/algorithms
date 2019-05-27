package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for i := range nums {
		idx := int(math.Abs(float64(nums[i])) - 1)
		if nums[idx] < 0 {
			result = append(result, idx+1)
		}
		nums[idx] = -nums[idx]
	}
	return result
}

func main() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(input))
}

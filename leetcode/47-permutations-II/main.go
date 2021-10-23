package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	helper(nums, []int{}, &result)
	return result
}

func helper(nums, subres []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, append([]int{}, subres...))
		return
	}
	for i := range nums {
		// If the prev number is the same, we can only use this number if prev is already used
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		helper(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(subres, nums[i]), result)
	}
}

func main() {
	input := []int{1, 3, 3, 3}
	fmt.Println(permuteUnique(input))
}

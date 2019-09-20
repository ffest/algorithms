package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	helper(nums, []int{}, &result)
	return result
}

func helper(nums, subres []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, subres)
		return
	}
	for i := range nums {
		numsCopy := append([]int{}, nums...)
		subresCopy := append([]int{}, subres...)
		helper(append(numsCopy[:i], numsCopy[i+1:]...), append(subresCopy, nums[i]), result)
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
}

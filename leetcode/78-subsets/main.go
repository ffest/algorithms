package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	result := make([][]int, 0)
	backtracking(nums, []int{}, &result)
	return result
}

func backtracking(nums, path []int, result *[][]int) {
	*result = append(*result, path)

	for i := 0; i < len(nums); i++ {
		// copy items
		nextPath := make([]int, 0, len(path)+1)
		nextPath = append(nextPath, path...)
		nextPath = append(nextPath, nums[i])
		backtracking(nums[i+1:], nextPath, result)
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(subsets(input))
}

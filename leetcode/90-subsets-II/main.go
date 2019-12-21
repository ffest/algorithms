package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	backtracking(nums, []int{}, &result)
	return result
}

func backtracking(nums, path []int, result *[][]int) {
	*result = append(*result, append([]int{}, path...))
	for i := 0; i < len(nums); {
		path = append(path, nums[i])
		backtracking(nums[i+1:], path, result)
		path = path[:len(path)-1]

		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

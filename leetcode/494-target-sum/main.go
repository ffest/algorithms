package main

import (
	"fmt"
)

// TODO: upd solution to n^2 dp
func findTargetSumWays(nums []int, S int) int {
	var count int
	helper(0, 0, S, nums, &count)
	return count
}

func helper(idx, sum, target int, nums []int, count *int) {
	if idx == len(nums) {
		if sum == target {
			*count++
		}
		return
	}

	helper(idx+1, sum+nums[idx], target, nums, count)
	helper(idx+1, sum-nums[idx], target, nums, count)
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}

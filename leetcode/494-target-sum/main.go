package main

import (
	"fmt"
)

// Slow 2^n solution
/*func findTargetSumWays(nums []int, S int) int {
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
}*/

func findTargetSumWays(nums []int, S int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+S)%2 == 1 || S > sum {
		return 0
	}
	target := (sum + S) / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for j := sum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
}

func main() {
	nums := []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	target := 0
	fmt.Println(findTargetSumWays(nums, target))
}

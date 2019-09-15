package main

import (
	"fmt"
)

// O(nlogn) solution
func lengthOfLIS(nums []int) int {
	var size int

	tails := make([]int, len(nums))
	for _, num := range nums {
		i, j := 0, size
		for i != j {
			mid := (i + j) / 2
			if tails[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		tails[i] = num
		if i == size {
			size++
		}
	}

	return size
}

// Naive DP O(n^2) solution
/*func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}*/

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

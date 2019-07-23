package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2

	dp := make([]bool, sum+1)
	dp[0] = true

	for _, num := range nums {
		for j := sum; j > 0; j-- {
			if j >= num && dp[j-num] {
				dp[j] = dp[j-num]
			}
		}
	}
	return dp[sum]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Print(canPartition(nums))
}

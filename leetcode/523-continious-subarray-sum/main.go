package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	if k == 0 {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == 0 && nums[i+1] == 0 {
				return true
			}
		}
		return false
	}

	currentSum := nums[0]
	sumsMap := make(map[int]int)
	sumsMap[0] = -1
	var mod int
	for i := 1; i < len(nums); i++ {
		mod = (currentSum + nums[i]) % k
		if _, ok := sumsMap[mod]; ok {
			return true
		}
		sumsMap[currentSum%k] = i - 1
		currentSum += nums[i]
	}
	return false
}

func main() {
	k := 2
	input := []int{1, 1}
	fmt.Println(checkSubarraySum(input, k))
}

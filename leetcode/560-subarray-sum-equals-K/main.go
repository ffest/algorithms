package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	var count int
	var sum int
	prefixSums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSums[i] = sum
	}
	cache := make(map[int]int)
	cache[0] = 1
	for i := 0; i < len(prefixSums); i++ {
		if _, ok := cache[prefixSums[i]-k]; ok {
			count += cache[prefixSums[i]-k]
		}
		cache[prefixSums[i]]++
	}
	return count
}

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(subarraySum(nums, 0))
}

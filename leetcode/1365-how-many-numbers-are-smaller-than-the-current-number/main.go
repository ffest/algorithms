package main

import (
	"fmt"
)

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))

	cache := [101]int{}
	for i := 0; i < len(nums); i++ {
		cache[nums[i]]++
	}
	for i := 1; i < len(cache); i++ {
		cache[i] += cache[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = 0
		} else {
			result[i] = cache[nums[i]-1]
		}
	}

	return result
}

func main() {
	nums := []int{8, 1, 2, 2, 3, 0}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

package main

import (
	"fmt"
)

func findLHS(nums []int) int {
	buckets := make(map[int]int)
	for i := range nums {
		buckets[nums[i]]++
	}

	var max int
	for i := 0; i < len(nums); i++ {
		if buckets[nums[i]+1] == 0 {
			continue
		}
		if buckets[nums[i]]+buckets[nums[i]+1] > max {
			max = buckets[nums[i]] + buckets[nums[i]+1]
		}
	}
	return max
}

func main() {
	input := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(input))
}

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 1<<31 - 1
	var left, right, sum int
	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-result)) {
				result = sum
			}
		}
	}
	return result
}

func main() {
	nums := []int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}
	target := 0
	fmt.Println(threeSumClosest(nums, target))
}

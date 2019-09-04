package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	findNSum(nums, target, 4, []int{}, &results)
	return results
}

func findNSum(nums []int, target, n int, result []int, results *[][]int) {
	if len(nums) < n || n < 2 {
		return
	}

	if n == 2 {
		left, right := 0, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				*results = append(*results, append(result, nums[left], nums[right]))
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	} else {
		for i := 0; i < len(nums)-n+1; i++ {
			if nums[i]*n > target || nums[len(nums)-1]*n < target {
				return
			} else if i == 0 || i > 0 && nums[i-1] != nums[i] {
				findNSum(nums[i+1:], target-nums[i], n-1, append(result, nums[i]), results)
			}
		}
	}
}

func main() {
	nums := []int{0, 0, 0, 0}
	target := 0
	fmt.Println(fourSum(nums, target))
}

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			left, right := i+1, len(nums)-1
			for left < right {
				if nums[left]+nums[right] > -nums[i] {
					right--
				} else if nums[left]+nums[right] < -nums[i] {
					left++
				} else {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						left--
					}
					left++
					right--
				}
			}
		}
	}
	return result
}

func main() {
	input := []int{-1, 0, 1, -1, 2, -4}
	fmt.Println(threeSum(input))
}

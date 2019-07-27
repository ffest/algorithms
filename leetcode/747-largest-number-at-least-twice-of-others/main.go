package main

import "fmt"

func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return len(nums) - 1
	}

	var first, second int
	if nums[0] > nums[1] {
		second = 1
	} else {
		first = 1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[first] {
			second = first
			first = i
		} else if nums[i] > nums[second] {
			second = i
		}
	}
	if nums[second] == 0 && nums[first] != 0 || nums[first]/nums[second] >= 2 {
		return first
	}
	return -1
}

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}

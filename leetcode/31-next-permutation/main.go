package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var idx = len(nums) - 1
	for idx > 0 {
		if nums[idx-1] < nums[idx] {
			break
		}
		idx--
	}
	// Need to reverse sort in any case
	reverseSort(nums, idx, len(nums)-1)
	if idx == 0 {
		return
	}

	for i := idx; i < len(nums); i++ {
		if nums[i] > nums[idx-1] {
			nums[idx-1], nums[i] = nums[i], nums[idx-1]
			return
		}
	}

}

func reverseSort(nums []int, start int, end int) {
	// Just swap them from the ends.
	for i, j := start, end; i < end+(start-end)/2; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{2, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

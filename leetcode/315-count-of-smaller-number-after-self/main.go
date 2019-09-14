package main

import (
	"fmt"
)

func countSmaller(nums []int) []int {
	idx := make([]int, 0, len(nums))
	for i := range nums {
		idx = append(idx, i)
	}
	inversions := make([]int, len(nums), len(nums))
	Merge(0, len(nums)-1, nums, idx, inversions)
	return inversions
}

func Merge(left, right int, nums, idx, inversions []int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	Merge(left, mid, nums, idx, inversions)
	Merge(mid+1, right, nums, idx, inversions)

	merge(left, right, nums, idx, inversions)
}

func merge(left, right int, nums, idx, inversions []int) {
	start, end := left, right
	mid := (left + right) / 2
	leftMid, rightMid := mid, mid+1

	newIdx := make([]int, 0)
	smaller := 0
	for left <= leftMid && rightMid <= right {
		if nums[idx[left]] <= nums[idx[rightMid]] {
			newIdx = append(newIdx, idx[left])
			inversions[idx[left]] += smaller
			left++
		} else {
			newIdx = append(newIdx, idx[rightMid])
			rightMid++
			smaller++
		}
	}

	for left <= leftMid {
		inversions[idx[left]] += smaller
		newIdx = append(newIdx, idx[left])
		left++
	}
	for rightMid <= right {
		newIdx = append(newIdx, idx[rightMid])
		rightMid++
	}

	copy(idx[start:end+1], newIdx)
}

func main() {
	//input := []int{5, 2, 6, 1}
	input := []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}
	fmt.Println(countSmaller(input))
}

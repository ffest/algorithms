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
	newIdx := make([]int, 0, len(idx))
	Merge(0, len(nums)-1, nums, idx, newIdx, inversions)
	return inversions
}

func Merge(left, right int, num, idx, newIdx, inversions []int) {
	if left >= right {
		return
	}

	middle := (left + right) / 2
	Merge(left, middle, num, idx, newIdx, inversions)
	Merge(middle+1, right, num, idx, newIdx, inversions)

	newIdx = newIdx[:0]
	merge(left, middle, middle+1, right, num, idx, newIdx, inversions)
}

func merge(left, leftMid, rightMid, right int, num, idx, newIdx, inversions []int) {
	start := left
	end := right
	smaller := 0
	for left <= leftMid && rightMid <= right {
		if num[idx[left]] > num[idx[rightMid]] {
			newIdx = append(newIdx, idx[rightMid])
			smaller++
			rightMid++
		} else {
			inversions[idx[left]] += smaller
			newIdx = append(newIdx, idx[left])
			left++
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

package main

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums)/2

	left := sortArray(nums[:middle])
	right := sortArray(nums[middle:])

	return merge(left, right)
}


func merge(left,right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}

		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}


func main() {
	nums := []int{5,2,3,1}
	fmt.Println(sortArray(nums))
}

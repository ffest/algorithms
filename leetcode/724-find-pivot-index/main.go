package main

import "fmt"

func pivotIndex(nums []int) int {
	var rightSum int
	for i := range nums {
		rightSum += nums[i]
	}
	var leftSum int
	for i := range nums {
		leftSum += nums[i]
		if leftSum == rightSum {
			return i
		}
		rightSum -= nums[i]
	}
	return -1
}

func main() {
	input := []int{5, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(input))
}

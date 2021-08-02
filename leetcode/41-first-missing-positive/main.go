package main

import "fmt"

func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] >= n {
			nums[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		nums[nums[i]%n] += n
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]/n == 0 {
			return i
		}
	}
	return n
}

func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	window, result := make([]int, 0), make([]int, 0)
	for i := range nums {
		for len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i > k-2 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

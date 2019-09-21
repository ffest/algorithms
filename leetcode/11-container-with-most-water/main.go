package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var current, result int
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			current = height[left] * (right - left)
			left++
		} else {
			current = height[right] * (right - left)
			right--
		}
		if current > result {
			result = current
		}
	}
	return result
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 9, 3, 7}
	fmt.Println(maxArea(input))
}

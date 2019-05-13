package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var result int
	left := 0
	right := len(height) - 1
	for left < right {
		if result < min(height[left], height[right])*(right-left) {
			result = min(height[left], height[right]) * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 9, 3, 7}
	fmt.Println(maxArea(input))
}

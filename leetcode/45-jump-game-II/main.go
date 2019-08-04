package main

import (
	"fmt"
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var jumps, curEnd, curFarthest int
	for i := 0; i < len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}

func main() {
	//input := []int{3, 2, 1, 3, 4, 5, 1, 1, 5, 3, 1, 5}
	input := []int{2, 3, 1, 1, 4}
	//input := []int{1, 0}
	fmt.Println(jump(input))
}

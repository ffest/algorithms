package main

import (
	"fmt"
)

func minSubArrayLen(s int, nums []int) int {
	var sum, min, head, tail int
	for tail = 0; tail < len(nums); tail++ {
		for sum < s && head < len(nums) {
			sum += nums[head]
			head++
		}
		if sum >= s && (min == 0 || head-tail < min) {
			min = head - tail
		}
		sum -= nums[tail]
	}
	return min
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}

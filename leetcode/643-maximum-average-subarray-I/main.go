package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var head, sum int
	for head < k {
		sum += nums[head]
		head++
	}
	globalSum := sum
	for ; head < len(nums); head++ {
		sum = sum + nums[head] - nums[head-k]
		if sum > globalSum {
			globalSum = sum
		}
	}
	return float64(globalSum) / float64(k)
}

func main() {
	input := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(input, k))
}

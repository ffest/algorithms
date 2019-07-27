package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var max, localMax int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			localMax++
		} else {
			if localMax > max {
				max = localMax
			}
			localMax = 0
		}
	}
	if localMax > max {
		max = localMax
	}
	return max
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

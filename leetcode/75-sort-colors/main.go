package main

import (
	"log"
)

func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	i := 0
	for i < right {
		if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			continue
		}
		i++
	}
	i = right
	for i > left {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			continue
		}
		i--
	}
}

func main() {
	input := []int{2, 1, 0, 2, 1, 1, 0, 1, 2, 2, 2, 0, 0}
	sortColors(input)
	log.Println(input)
}

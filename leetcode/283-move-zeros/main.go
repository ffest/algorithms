package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var current int
	for i := range nums {
		if nums[i] != 0 {
			nums[current] = nums[i]
			current++
		}
	}
	for i := current; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

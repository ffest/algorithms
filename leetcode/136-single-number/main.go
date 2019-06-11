package main

import "fmt"

func singleNumber(nums []int) int {
	var result int
	for i := range nums {
		result ^= nums[i]
	}
	return result
}

func main() {
	input := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(input))
}

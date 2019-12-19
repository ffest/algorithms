package main

import (
	"fmt"
	"log"
)

func removeDuplicates(nums []int) int {
	var writeIdx, prev, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == prev && count >= 2 {
			continue
		}
		if nums[i] != prev {
			count = 0
		}
		nums[writeIdx] = nums[i]
		writeIdx++
		count++
		prev = nums[i]
	}
	return writeIdx
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	log.Print(nums)
}

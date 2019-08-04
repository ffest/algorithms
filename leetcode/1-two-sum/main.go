package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := range nums {
		if idx, ok := cache[nums[i]]; ok {
			return []int{idx, i}
		}
		cache[target-nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

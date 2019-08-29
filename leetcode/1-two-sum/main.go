package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, n := range nums {
		if c, ok := cache[n]; ok {
			return []int{c, i}
		}
		cache[target-n] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22
	fmt.Println(twoSum(nums, target))
}

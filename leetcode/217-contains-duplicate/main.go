package main

import "fmt"

func containsDuplicate(nums []int) bool {
	cache := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := cache[n]; ok {
			return true
		}
		cache[n] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

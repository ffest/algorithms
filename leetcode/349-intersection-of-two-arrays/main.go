package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	for _, n := range nums1 {
		cache[n] = 0
	}
	for _, n := range nums2 {
		if _, ok := cache[n]; ok {
			cache[n]++
		}
	}
	result := make([]int, 0)
	for k, v := range cache {
		if v == 0 {
			continue
		}
		result = append(result, k)
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

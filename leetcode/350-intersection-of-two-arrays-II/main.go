package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	firstHash := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		firstHash[num]++
	}

	var result []int
	for _, num := range nums2 {
		if q, ok := firstHash[num]; ok && q > 0 {
			result = append(result, num)
			firstHash[num]--
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}

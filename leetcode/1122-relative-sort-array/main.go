package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	cache := make(map[int]int)
	for _, num := range arr1 {
		cache[num]++
	}

	result := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		count := cache[num]
		for i := 0; i < count; i++ {
			result = append(result, num)
		}
		delete(cache, num)
	}

	left := make([]int, 0, len(cache))
	for k, v := range cache {
		for i := 0; i < v; i++ {
			left = append(left, k)
		}
	}
	sort.Ints(left)
	result = append(result, left...)

	return result
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	cache := make(map[string]int)
	for i, l := range list1 {
		cache[l] = i
	}
	var result []string
	leastIndexSum := 1<<31 - 1

	for i1, l := range list2 {
		if i2, ok := cache[l]; ok {
			indexSum := i1 + i2
			if indexSum < leastIndexSum {
				leastIndexSum = indexSum
				result = []string{l}
			} else if indexSum == leastIndexSum {
				result = append(result, l)
			}
		}
	}

	return result
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(list1, list2))
}

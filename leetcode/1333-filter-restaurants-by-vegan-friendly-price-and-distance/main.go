package main

import (
	"fmt"
	"sort"
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := make([][]int, 0)
	for _, rest := range restaurants {
		if rest[3] > maxPrice || rest[4] > maxDistance || (veganFriendly == 1 && rest[2] == 0) {
			continue
		}
		filtered = append(filtered, rest)
	}
	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i][1] == filtered[j][1] {
			return filtered[i][0] > filtered[j][0]
		}
		return filtered[i][1] > filtered[j][1]
	})
	result := make([]int, len(filtered))
	for i := range filtered {
		result[i] = filtered[i][0]
	}

	return result
}

func main() {
	restourants := [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	fmt.Println(filterRestaurants(restourants, 1, 50, 10))
}

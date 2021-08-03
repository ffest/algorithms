package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	best, count := 0, 0
	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}
		current := num
		for {
			if _, ok := set[current]; ok {
				current++
				count++
			} else {
				if best < count {
					best = count
				}
				count = 0
				break
			}
		}
	}
	return best
}

func main() {
	nums := []int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}
	fmt.Println(longestConsecutive(nums))
}

package main

import "fmt"

func minCostToMoveChips(chips []int) int {
	if len(chips) < 2 {
		return 0
	}
	var odd, even int
	for _, chip := range chips {
		if chip%2 == 1 {
			odd++
		} else {
			even++
		}
	}

	if odd < even {
		return odd
	}
	return even
}

func main() {
	chips := []int{1, 2, 3}
	fmt.Println(minCostToMoveChips(chips))
}

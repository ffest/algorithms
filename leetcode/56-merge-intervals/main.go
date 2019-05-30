package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	merged := make([][]int, 0, len(intervals))

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var current []int
	for i := range intervals {
		switch {
		case current == nil:
			current = intervals[i]
		case intervals[i][0] > current[1]:
			merged = append(merged, current)
			current = intervals[i]
		case intervals[i][1] > current[1]:
			current[1] = intervals[i][1]
		}
	}
	if current != nil {
		merged = append(merged, current)
	}
	return merged
}

func main() {
	intervals := [][]int{
		{1, 4},
		{0, 0},
	}

	fmt.Println(merge(intervals))
}

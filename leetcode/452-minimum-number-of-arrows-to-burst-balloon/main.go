package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	arrows := 1
	position := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= position {
			continue
		}
		arrows++
		position = points[i][1]
	}
	return arrows
}

func main() {
	points := [][]int{
		{3, 9},
		{7, 12},
		{3, 8},
		{6, 8},
		{9, 10},
		{2, 9},
		{0, 9},
		{3, 9},
		{0, 6},
		{2, 8},
	}
	fmt.Println(findMinArrowShots(points))
}

package main

import (
	"fmt"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var count int
	topBottomMax := make([]int, 0, len(grid))
	leftRightMax := make([]int, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		lrMax := 0
		tbMax := 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > lrMax {
				lrMax = grid[i][j]
			}
			if grid[j][i] > tbMax {
				tbMax = grid[j][i]
			}
		}
		leftRightMax = append(leftRightMax, lrMax)
		topBottomMax = append(topBottomMax, tbMax)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if topBottomMax[j] > leftRightMax[i] {
				count += leftRightMax[i] - grid[i][j]
			} else {
				count += topBottomMax[j] - grid[i][j]
			}
		}
	}

	return count
}

func main() {
	var grid [][]int

	// example
	grid = append(grid, []int{3, 0, 8, 4})
	grid = append(grid, []int{2, 4, 5, 7})
	grid = append(grid, []int{9, 2, 6, 3})
	grid = append(grid, []int{0, 3, 1, 0})

	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

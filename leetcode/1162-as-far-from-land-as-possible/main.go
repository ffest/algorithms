package main

import (
	"fmt"
	"math"
)

func maxDistance(grid [][]int) int {
	earthes := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				earthes = append(earthes, [2]int{i, j})
			}
		}
	}
	if len(earthes) == 0 || len(earthes) == len(grid)*len(grid[0]) {
		return -1
	}
	var maxDistance int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				continue
			}

			currentMin := math.MaxInt32
			for _, earth := range earthes {
				current := int(math.Abs(float64(earth[0]-i)) + math.Abs(float64(earth[1]-j)))
				if current < currentMin {
					currentMin = current
				}
			}
			if currentMin > maxDistance {
				maxDistance = currentMin
			}
		}
	}

	return maxDistance
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(maxDistance(grid))
}

package main

import (
	"fmt"
)

func maxDistance(grid [][]int) int {
	queue := make([][2]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i + 1, j}, [2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i, j - 1})
			}
		}
	}
	if len(queue) == 0 || len(grid)*len(grid[0]) == len(queue)/4 {
		return -1
	}
	var steps int
	for len(queue) > 0 {
		steps++
		qSize := len(queue)
		for k := 0; k < qSize; k++ {
			current := queue[0]
			queue = queue[1:]
			i, j := current[0], current[1]
			if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] == 0 {
				grid[i][j] = steps
				queue = append(queue, [2]int{i + 1, j}, [2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i, j - 1})
			}
		}
	}
	return steps - 1
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}
	fmt.Println(maxDistance(grid))
}

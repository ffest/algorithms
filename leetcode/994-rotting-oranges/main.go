package main

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	var total, rotted int
	n, m := len(grid), len(grid[0])
	queue := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 || grid[i][j] == 2 {
				total++
				if grid[i][j] == 2 {
					queue = append(queue, [2]int{i, j})
				}
			}
		}
	}

	if total == 0 {
		return 0
	}

	var timer int
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			rotted++
			current := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				x, y := current[0]+d[0], current[1]+d[1]
				if x < 0 || y < 0 || x == n || y == m || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				queue = append(queue, [2]int{x, y})
			}
		}
		timer++
	}

	if rotted == total {
		return timer - 1
	}
	return -1
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}

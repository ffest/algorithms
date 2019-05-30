package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dist[i][j] = grid[i][j]
			} else if i == 0 {
				dist[i][j] = grid[i][j] + dist[i][j-1]
			} else if j == 0 {
				dist[i][j] = grid[i][j] + dist[i-1][j]
			} else {
				dist[i][j] = grid[i][j] + min(dist[i-1][j], dist[i][j-1])
			}
		}
	}

	return dist[n-1][m-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	grid := [][]int{
		{1, 3, 1, 4},
		{2, 8, 1, 2},
		{6, 2, 12, 4},
		{8, 3, 1, 7},
	}
	fmt.Println(minPathSum(grid))
}

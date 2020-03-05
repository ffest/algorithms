package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid)
				count++
			}
		}
	}
	return count
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j]++
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '1'},
	}

	fmt.Println(numIslands(grid))
}

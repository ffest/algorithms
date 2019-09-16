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
			if grid[i][j] != '1' {
				continue
			}

			count++
			helper(grid, i, j)
		}
	}
	return count
}

func helper(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j]++
	helper(grid, i+1, j)
	helper(grid, i-1, j)
	helper(grid, i, j+1)
	helper(grid, i, j-1)
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

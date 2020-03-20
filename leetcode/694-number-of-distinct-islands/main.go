package main

import (
	"fmt"
)

func numDistinctIslands(grid [][]int) int {
	var count int
	cache := map[string]struct{}{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			var path string
			dfs(grid, i, j, &path, "s")
			if _, ok := cache[path]; !ok {
				count++
				cache[path] = struct{}{}
			}
		}
	}
	return count
}

func dfs(grid [][]int, i, j int, path *string, direction string) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
		return
	}
	*path += direction
	grid[i][j] = 0
	dfs(grid, i+1, j, path, "d")
	dfs(grid, i-1, j, path, "u")
	dfs(grid, i, j+1, path, "r")
	dfs(grid, i, j-1, path, "l")
	*path += "b"
}

func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numDistinctIslands(grid))
}

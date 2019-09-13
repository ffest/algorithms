package main

import (
	"fmt"
)

func uniquePathsIII(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var obstacles int
	start := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				start = [2]int{i, j}
			} else if grid[i][j] == -1 {
				obstacles++
			}
		}
	}
	visited := make(map[[2]int]struct{})
	var result int
	dfs(len(grid)*len(grid[0])-obstacles, grid, visited, start, &result)
	return result
}

func dfs(needToVisit int, grid [][]int, visited map[[2]int]struct{}, point [2]int, result *int) {
	if point[0] < 0 || point[0] == len(grid) || point[1] < 0 || point[1] == len(grid[0]) || grid[point[0]][point[1]] == -1 {
		return
	}
	if _, ok := visited[point]; ok {
		return
	}
	if grid[point[0]][point[1]] == 2 && len(visited) == needToVisit-1 {
		*result++
		return
	}
	visited[point] = struct{}{}
	dfs(needToVisit, grid, visited, [2]int{point[0] + 1, point[1]}, result)
	dfs(needToVisit, grid, visited, [2]int{point[0] - 1, point[1]}, result)
	dfs(needToVisit, grid, visited, [2]int{point[0], point[1] + 1}, result)
	dfs(needToVisit, grid, visited, [2]int{point[0], point[1] - 1}, result)
	delete(visited, point)
}

func main() {
	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
	fmt.Println(uniquePathsIII(grid))
}

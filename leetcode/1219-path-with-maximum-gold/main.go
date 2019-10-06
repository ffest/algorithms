package main

import "fmt"

func getMaximumGold(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var max int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			sum := 0
			dfs(grid, i, j, &sum, &max)
		}
	}

	return max
}

func dfs(grid [][]int, i, j int, sum, max *int) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] <= 0 {
		return
	}
	*sum += grid[i][j]
	if *max < *sum {
		*max = *sum
	}
	grid[i][j] *= -1
	dfs(grid, i+1, j, sum, max)
	dfs(grid, i-1, j, sum, max)
	dfs(grid, i, j+1, sum, max)
	dfs(grid, i, j-1, sum, max)
	grid[i][j] *= -1
	*sum -= grid[i][j]
}

func main() {
	grid := [][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	}
	fmt.Println(getMaximumGold(grid))
}

package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	return 0
}

func main() {
	var grid [][]int

	// example
	grid = append(grid, []int{1, 0, 0, 0})
	grid = append(grid, []int{0, 0, 0, 0})
	grid = append(grid, []int{0, 0, 2, -1})

	fmt.Println(uniquePathsIII(grid))
}

package main

import (
	"fmt"
)

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	n, m := len(matrix), len(matrix[0])
	coordinates := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				coordinates = append(coordinates, [2]int{i, j})
				continue
			}
			matrix[i][j] = 1<<31 - 1
		}
	}

	ways := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(coordinates) > 0 {
		current := coordinates[0]
		coordinates = coordinates[1:]
		currentVal := matrix[current[0]][current[1]]

		for _, way := range ways {
			i, j := current[0]+way[0], current[1]+way[1]
			if i < 0 || j < 0 || i == n || j == m || matrix[i][j] == 0 || matrix[i][j] < currentVal+1 {
				continue
			}
			matrix[i][j] = currentVal + 1
			coordinates = append(coordinates, [2]int{i, j})
		}
	}

	return matrix
}

func main() {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
		{1, 1, 0},
	}
	fmt.Println(updateMatrix(matrix))
}

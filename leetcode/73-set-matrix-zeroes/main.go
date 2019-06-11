package main

import (
	"fmt"
)

func setZeroesAdditionalSpace(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n := len(matrix)
	m := len(matrix[0])
	rows := make(map[int]struct{}, n)
	columns := make(map[int]struct{}, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				columns[j] = struct{}{}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, okRow := rows[i]
			_, okColumn := columns[j]
			if okRow || okColumn {
				matrix[i][j] = 0
			}
		}
	}
	return
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n := len(matrix)
	m := len(matrix[0])
	var fr, fc bool

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if fc {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	input := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(input)
	fmt.Println(input)
}

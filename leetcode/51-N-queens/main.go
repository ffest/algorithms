package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	columns := make([]bool, n) // columns   |
	d1 := make([]bool, 2*n)    // diagonals \
	d2 := make([]bool, 2*n)    // diagonals /
	candidate := make([][]rune, n)
	for i := range candidate {
		candidate[i] = make([]rune, n)
		for j := range candidate[i] {
			candidate[i][j] = '.'
		}
	}

	backtracking(candidate, 0, n, columns, d1, d2, &result)
	return result
}

func backtracking(candidate [][]rune, row, n int, columns, d1, d2 []bool, res *[][]string) {
	if row == n {
		var arr []string
		for _, v := range candidate {
			arr = append(arr, string(v))
		}
		*res = append(*res, arr)
		return
	}
	for i := 0; i < n; i++ {
		idx1 := n - i + row
		idx2 := row + i
		if d1[idx1] || d2[idx2] || columns[i] {
			continue
		}
		d1[idx1], d2[idx2], columns[i] = true, true, true
		candidate[row][i] = 'Q'
		backtracking(candidate, row+1, n, columns, d1, d2, res)
		d1[idx1], d2[idx2], columns[i] = false, false, false
		candidate[row][i] = '.'
	}
}

func main() {
	fmt.Println(solveNQueens(4))
}

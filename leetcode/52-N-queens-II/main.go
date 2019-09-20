package main

import (
	"fmt"
)

func totalNQueens(n int) int {
	var count int
	columns := make([]bool, n)
	d1 := make([]bool, 2*n)
	d2 := make([]bool, 2*n)
	backtracking(0, n, columns, d1, d2, &count)
	return count
}

func backtracking(currentRow, n int, columns, d1, d2 []bool, count *int) {
	if currentRow == n {
		*count++
		return
	}

	for i := 0; i < n; i++ {
		idx1, idx2 := i+currentRow, n-i+currentRow
		if columns[i] || d1[idx1] || d2[idx2] {
			continue
		}

		columns[i], d1[idx1], d2[idx2] = true, true, true
		backtracking(currentRow+1, n, columns, d1, d2, count)
		columns[i], d1[idx1], d2[idx2] = false, false, false
	}
}

func main() {
	fmt.Println(totalNQueens(2))
}

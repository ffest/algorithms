package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	prevRow := []int{1, 1}

	var currentRow []int
	for i := 2; i <= rowIndex; i++ {
		currentRow = make([]int, i+1)
		currentRow[0], currentRow[len(currentRow)-1] = 1, 1
		for j := 1; j < len(currentRow)-1; j++ {
			currentRow[j] = prevRow[j-1] + prevRow[j]
		}
		prevRow = currentRow
	}
	return currentRow
}

func main() {
	input := 3
	fmt.Println(getRow(input))
}

package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	triange := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		triange = append(triange, make([]int, i+1))
		triange[i][0], triange[i][len(triange[i])-1] = 1, 1
		for j := 1; j < len(triange[i])-1; j++ {
			triange[i][j] = triange[i-1][j-1] + triange[i-1][j]
		}
	}
	return triange
}

func main() {
	input := 8
	fmt.Println(generate(input))
}

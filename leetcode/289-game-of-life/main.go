package main

import "fmt"

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := countNeighbours(board, i, j)
			if board[i][j] == 1 && (count == 2 || count == 3) {
				board[i][j] = 3
			}
			if board[i][j] == 0 && count == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

func countNeighbours(board [][]int, i, j int) int {
	var count int
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k < 0 || k == len(board) || l < 0 || l == len(board[0]) || (k == i && l == j) {
				continue
			}
			count += board[k][l] & 1
		}
	}
	return count
}

func main() {
	pes := 1
	fmt.Println(pes >> 1)
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}

package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	blocksCache := make([]map[byte]struct{}, 9)
	rowCache := make([]map[byte]struct{}, 9)
	columnCache := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		blocksCache[i] = make(map[byte]struct{})
		rowCache[i] = make(map[byte]struct{})
		columnCache[i] = make(map[byte]struct{})
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			blockKey := (i/3)*3 + j/3
			if _, ok := blocksCache[blockKey][board[i][j]]; ok {
				return false
			}
			if _, ok := rowCache[i][board[i][j]]; ok {
				return false
			}
			if _, ok := columnCache[j][board[i][j]]; ok {
				return false
			}
			blocksCache[blockKey][board[i][j]] = struct{}{}
			rowCache[i][board[i][j]] = struct{}{}
			columnCache[j][board[i][j]] = struct{}{}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

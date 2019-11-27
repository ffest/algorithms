package main

import "fmt"

func countBattleships(board [][]byte) int {
	var count int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				count++
			}
		}
	}
	return count
}

func main() {
	board := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println(countBattleships(board))
}

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, word, board) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j int, word string, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = '#'
	res := dfs(i+1, j, word[1:], board) || dfs(i-1, j, word[1:], board) ||
		dfs(i, j+1, word[1:], board) || dfs(i, j-1, word[1:], board)
	board[i][j] = tmp
	return res
}

func main() {
	/*board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}*/
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	fmt.Println(exist(board, "AAB"))
}

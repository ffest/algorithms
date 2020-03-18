package main

import (
	"fmt"
)

type TicTacToe struct {
	n            int
	rows         []int
	columns      []int
	diagonal     int
	antiDiagonal int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{
		n:       n,
		rows:    make([]int, n),
		columns: make([]int, n),
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (tt *TicTacToe) Move(row int, col int, player int) int {
	toAdd := 1
	if player == 2 {
		toAdd *= -1
	}

	tt.rows[row] += toAdd
	tt.columns[col] += toAdd
	if row == col {
		tt.diagonal += toAdd
	}
	if tt.n-row-col-1 == 0 {
		tt.antiDiagonal += toAdd
	}

	if abs(tt.rows[row]) == tt.n || abs(tt.columns[col]) == tt.n ||
		abs(tt.diagonal) == tt.n || abs(tt.antiDiagonal) == tt.n {
		return player
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	tt := Constructor(2)
	fmt.Println(tt.Move(0, 0, 2))
	fmt.Println(tt.Move(0, 1, 1))
	fmt.Println(tt.Move(1, 1, 2))
}

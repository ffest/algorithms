package main

import (
	"fmt"
)

func minFallingPathSum(A [][]int) int {
	if len(A) == 1 {
		return A[0][0]
	}
	prev, minPathSum := A[0], 1<<31-1
	for i := 1; i < len(A); i++ {
		cur := A[i]
		for j := 0; j < len(A); j++ {
			prevMin := prev[j]
			if j != 0 {
				prevMin = min(prevMin, prev[j-1])
			}
			if j != len(A)-1 {
				prevMin = min(prevMin, prev[j+1])
			}
			cur[j] += prevMin
			if i == len(A)-1 && minPathSum > cur[j] {
				minPathSum = cur[j]
			}
		}
		prev = cur
	}
	return minPathSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	A := [][]int{
		{10, -98, 44},
		{-20, 65, 34},
		{-100, -1, 74},
	}
	fmt.Println(minFallingPathSum(A))
}

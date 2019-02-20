package main

import "fmt"

func minDeletionSize(A []string) int {
	stringLen := len(A)
	cn := 0
	for i := 0; i < len(A[0]); i++ {
		for k := 0; k < stringLen-1; k++ {
			if A[k][i] > A[k+1][i] {
				cn++
				break
			}
		}
	}
	return cn

}

func main() {
	input := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(input))
}

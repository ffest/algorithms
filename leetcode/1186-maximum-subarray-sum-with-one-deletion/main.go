package main

import (
	"fmt"
)

func maximumSum(arr []int) int {
	result := -1 << 31

	var maxNoDeletion, maxOneDeletion int
	maxNoDeletion = arr[0]
	for i := 1; i < len(arr); i++ {
		maxOneDeletion = max(maxNoDeletion, maxOneDeletion+arr[i])
		maxNoDeletion = max(maxNoDeletion+arr[i], arr[i])
		result = max(result, maxOneDeletion)
		result = max(result, maxNoDeletion)
	}
	return max(result, maxNoDeletion)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, -2, 0, 3}
	fmt.Println(maximumSum(arr))
}

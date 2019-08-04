package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtracking([]int{}, candidates, 0, target, &result)
	return result
}

func backtracking(subset, candidates []int, sum, target int, result *[][]int) {
	if sum == target {
		*result = append(*result, subset)
		return
	} else if sum > target {
		return
	}

	for i := range candidates {
		subsetCopy := make([]int, 0, len(subset)+1)
		subsetCopy = append(subsetCopy, subset...)
		backtracking(append(subsetCopy, candidates[i]), candidates[i:], sum+candidates[i], target, result)
	}
}

func main() {
	input := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(input, target))
}

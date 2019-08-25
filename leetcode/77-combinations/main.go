package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	helper(1, n, k, []int{}, &result)
	return result
}

func helper(pointer, n, k int, current []int,  result *[][]int) {
	if len(current) == k {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := pointer; i <= n; i++ {
		helper(i+1, n, k, append(current, i), result)
	}
}

func main() {
	fmt.Println(combine(4,2))
}

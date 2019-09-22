package main

import "fmt"

func subarraysWithKDistinct(A []int, K int) int {
	return atMostK(A, K) - atMostK(A, K-1)
}

func atMostK(A []int, k int) int {
	var left, result int
	count := [20001]int{}
	for right := 0; right < len(A); right++ {
		if count[A[right]] == 0 {
			k--
		}
		count[A[right]]++
		for k < 0 {
			count[A[left]]--
			if count[A[left]] == 0 {
				k++
			}
			left++
		}
		result += right - left + 1
	}
	return result
}

func main() {
	A := []int{1, 2, 1, 2, 3}
	fmt.Println(subarraysWithKDistinct(A, 2))
}

package main

import (
	"fmt"
)

func numSubarraysWithSum(A []int, S int) int {
	cache := make(map[int]int)
	cache[0] = 1
	var prefixSum, count int
	for _, num := range A {
		prefixSum += num
		count += cache[prefixSum-S]
		cache[prefixSum]++
	}
	return count
}

/*func numSubarraysWithSum(A []int, S int) int {
	return atMostK(A, S) - atMostK(A, S-1)
}

func atMostK(A []int, k int) int {
	if k < 0 {
		return 0
	}
	var i, j, count int
	for j = 0; j < len(A); j++ {
		k -= A[j]
		for k < 0 {
			k += A[i]
			i++
		}
		count += j - i + 1
	}
	return count
}*/

func main() {
	A := []int{1, 0, 1, 1, 0, 0, 1, 0, 1}
	fmt.Println(numSubarraysWithSum(A, 2))
}

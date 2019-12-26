package main

import (
	"fmt"
)

func kConcatenationMaxSum(arr []int, k int) int {
	maxSumSub := maxSumSubarray(arr)
	if k == 1 {
		return maxSumSub
	}
	var sum int
	for i := range arr {
		sum = (sum + arr[i]) % 1000000007
	}
	prefixSum, suffixSum := prefixSum(arr), suffixSum(arr)
	if sum > 0 {
		return max((sum*(k-2)+prefixSum+suffixSum)%1000000007, sum*k%1000000007)
	} else {
		return max(maxSumSub, (prefixSum+suffixSum)%1000000007)
	}
}

func maxSumSubarray(arr []int) int {
	var sum int
	maximum := -1 << 31
	for i := range arr {
		sum = max(sum+arr[i], arr[i])
		maximum = max(sum, maximum) % 1000000007
	}
	if maximum < 0 {
		return 0
	}
	return maximum % 1000000007
}

func prefixSum(arr []int) int {
	var sum int
	maximum := -1 << 31
	for i := 0; i < len(arr); i++ {
		sum = (sum + arr[i]) % 1000000007
		maximum = max(maximum, sum)
	}
	return maximum
}

func suffixSum(arr []int) int {
	var sum int
	maximum := -1 << 31
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		maximum = max(maximum, sum)
	}
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 2}
	k := 1
	fmt.Println(kConcatenationMaxSum(arr, k))
}

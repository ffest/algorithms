package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	result := make([]string, 0)

	numbers := make([]int, 0, n)
	factorials := make([]int, n+1)
	factorials[0] = 1
	for i := 1; i <= n; i++ {
		factorials[i] = factorials[i-1] * i
		numbers = append(numbers, i)
	}

	k--
	for i := 1; i <= n; i++ {
		idx := k / factorials[n-i]
		result = append(result, strconv.Itoa(numbers[idx]))
		numbers = append(numbers[:idx], numbers[idx+1:]...)
		k -= idx * factorials[n-i]
	}

	return strings.Join(result, "")
}

func main() {
	fmt.Println(getPermutation(4, 9))
}

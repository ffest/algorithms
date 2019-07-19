package main

import (
	"fmt"
)

// TODO: remind it again
func superEggDrop(K int, N int) int {
	var i int
	dp := make([]int, K+1)
	for i = 0; dp[K] < N; i++ {
		for j := K; j > 0; j-- {
			dp[j] += dp[j-1] + 1
		}
	}
	return i
}

func main() {
	k := 3
	n := 14
	fmt.Print(superEggDrop(k, n))
}

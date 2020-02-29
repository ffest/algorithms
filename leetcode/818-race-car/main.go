package main

import (
	"fmt"
	"math"
)

var dp [10001]int

func racecar(target int) int {
	if dp[target] > 0 {
		return dp[target]
	}
	n := int(math.Floor(math.Log2(float64(target)))) + 1
	if 1<<n == target+1 {
		dp[target] = n
	} else {
		dp[target] = racecar(1<<n-1-target) + n + 1
		for m := 0; m < n-1; m++ {
			dp[target] = min(dp[target], racecar(target-(1<<(n-1))+(1<<m))+n+m+1)
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(racecar(987))
}

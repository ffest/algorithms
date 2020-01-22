package main

import "fmt"

var mod = 1000000007

/*func knightDialer(N int) int {
	cache := make([][4][3]int, N+1)
	var result int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			result = (result + paths(cache, i, j, N)) % mod
		}
	}
	return result
}

func paths(cache [][4][3]int, i, j, n int) int {
	if i < 0 || j < 0 || i >= 4 || j >= 3 || (i == 3 && j != 1) {
		return 0
	}
	if n == 1 {
		return 1
	}
	if cache[n][i][j] > 0 {
		return cache[n][i][j]
	}
	cache[n][i][j] = paths(cache, i-2, j-1, n-1)%mod +
		paths(cache, i-1, j-2, n-1)%mod +
		paths(cache, i-2, j+1, n-1)%mod +
		paths(cache, i-1, j+2, n-1)%mod +
		paths(cache, i+2, j-1, n-1)%mod +
		paths(cache, i+1, j-2, n-1)%mod +
		paths(cache, i+2, j+1, n-1)%mod +
		paths(cache, i+1, j+2, n-1)%mod
	return cache[n][i][j]
}*/

func knightDialer(N int) int {
	if N == 1 {
		return 10
	}
	var hops = [][]int{
		{4, 6},    // from 0
		{6, 8},    // from 1
		{7, 9},    // from 2
		{4, 8},    // from 3
		{3, 9, 0}, // from 4
		{},        // from 5
		{0, 1, 7}, // from 6
		{2, 6},    // from 7
		{1, 3},    // from 8
		{2, 4},    // from 9
	}
	var sum int
	cache := [5000][10]int{}
	for i := 0; i < 10; i++ {
		sum = (sum + helper(N-1, i, hops, &cache)) % mod
	}
	return sum
}

func helper(n, digit int, hops [][]int, cache *[5000][10]int) int {
	if (*cache)[n][digit] > 0 {
		return (*cache)[n][digit]
	}
	if n == 1 {
		(*cache)[n][digit] = len(hops[digit])
		return (*cache)[n][digit]
	}

	for i := 0; i < len(hops[digit]); i++ {
		(*cache)[n][digit] = ((*cache)[n][digit] + helper(n-1, hops[digit][i], hops, cache)) % mod
	}
	return (*cache)[n][digit]
}

func main() {
	fmt.Println(knightDialer(1))
}

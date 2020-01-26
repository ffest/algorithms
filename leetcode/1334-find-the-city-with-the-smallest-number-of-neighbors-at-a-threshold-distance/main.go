package main

import (
	"fmt"
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dists := make([][]int, n)
	for i := range dists {
		dists[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				dists[i][j] = math.MaxInt32
			}
		}
	}
	for _, edge := range edges {
		dists[edge[0]][edge[1]] = edge[2]
		dists[edge[1]][edge[0]] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				} else if dists[i][k]+dists[k][j] < dists[i][j] {
					dists[i][j] = dists[i][k] + dists[k][j]
				}
			}
		}
	}
	var result int
	min := n
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if dists[i][j] <= distanceThreshold {
				count++
			}
		}
		if count <= min {
			min = count
			result = i
		}
	}
	return result
}

func main() {
	n := 5
	edges := [][]int{
		{0, 1, 2},
		{0, 4, 8},
		{1, 2, 3},
		{1, 4, 2},
		{2, 3, 1},
		{3, 4, 1},
	}
	distanceThreshold := 2
	fmt.Println(findTheCity(n, edges, distanceThreshold))
}

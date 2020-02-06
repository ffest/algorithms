package main

import (
	"fmt"
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	lengths := [2][]int{}
	lengths[0] = make([]int, n)
	lengths[1] = make([]int, n)
	for i := 1; i < n; i++ {
		lengths[0][i] = 2 * n
		lengths[1][i] = 2 * n
	}
	redCache, blueCache := make(map[int][]int), make(map[int][]int)
	for _, edge := range redEdges {
		redCache[edge[0]] = append(redCache[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueCache[edge[0]] = append(blueCache[edge[0]], edge[1])
	}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	queue = append(queue, [2]int{0, 1})
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			point := queue[0]
			queue = queue[1:]
			switch point[1] {
			// Last color was red
			case 0:
				for i := 0; i < len(blueCache[point[0]]); i++ {
					if lengths[1][blueCache[point[0]][i]] == 2*n {
						lengths[1][blueCache[point[0]][i]] = 1 + lengths[0][point[0]]
						queue = append(queue, [2]int{blueCache[point[0]][i], 1})
					}
				}
			// Last color was blue
			case 1:
				for i := 0; i < len(redCache[point[0]]); i++ {
					if lengths[0][redCache[point[0]][i]] == 2*n {
						lengths[0][redCache[point[0]][i]] = 1 + lengths[1][point[0]]
						queue = append(queue, [2]int{redCache[point[0]][i], 0})
					}
				}
			}
		}
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if lengths[0][i] < lengths[1][i] {
			result[i] = lengths[0][i]
		} else {
			result[i] = lengths[1][i]
		}
		if result[i] == 2*n {
			result[i] = -1
		}
	}
	return result
}

func main() {
	redEdges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	blueEdges := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	fmt.Println(shortestAlternatingPaths(5, redEdges, blueEdges))
}

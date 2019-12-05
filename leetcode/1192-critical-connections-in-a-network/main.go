package main

import (
	"fmt"
)

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], connection[1])
		graph[connection[1]] = append(graph[connection[1]], connection[0])
	}

	var dfs func(int, int)
	visited := make(map[int]int)
	var timer int

	critical := make([][]int, 0)
	dfs = func(v int, p int) {
		visited[v] = timer
		currentTime := timer
		timer++

		for i := 0; i < len(graph[v]); i++ {
			to := graph[v][i]
			if to == p {
				continue
			}

			if _, ok := visited[to]; !ok {
				dfs(to, v)
			}
			visited[v] = min(visited[to], visited[v])
			if currentTime < visited[to] {
				critical = append(critical, []int{v, to})
			}
		}
	}
	dfs(0, 0)

	return critical
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 4
	connections := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{1, 3},
	}
	fmt.Println(criticalConnections(n, connections))
}

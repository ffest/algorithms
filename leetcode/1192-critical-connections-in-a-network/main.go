package main

import (
	"fmt"
)

// Tarjan algorithm
/*func criticalConnections(n int, connections [][]int) [][]int {
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
}*/

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], connection[1])
		graph[connection[1]] = append(graph[connection[1]], connection[0])
	}

	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		ranks[i] = -2
	}
	cyclic := make(map[[2]int]struct{})
	var dfs func(node, depth int) int

	dfs = func(node, depth int) int {
		if ranks[node] >= 0 {
			return ranks[node]
		}
		ranks[node] = depth
		minBackDepth := n

		for _, next := range graph[node] {
			if ranks[next] == depth-1 {
				continue
			}
			backDepth := dfs(next, depth+1)
			if backDepth <= depth {
				cyclic[[2]int{node, next}] = struct{}{}
			}
			minBackDepth = min(minBackDepth, backDepth)
		}
		return minBackDepth
	}

	dfs(0, 0)
	result := make([][]int, 0)
	for _, connection := range connections {
		if _, ok := cyclic[[2]int{connection[0], connection[1]}]; ok {
			continue
		}
		if _, ok := cyclic[[2]int{connection[1], connection[0]}]; ok {
			continue
		}
		result = append(result, connection)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 5
	connections := [][]int{
		{1, 0},
		{2, 0},
		{3, 2},
		{4, 2},
		{4, 3},
		{3, 0},
		{4, 0},
	}
	fmt.Println(criticalConnections(n, connections))
}

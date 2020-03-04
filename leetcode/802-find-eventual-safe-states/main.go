package main

import (
	"fmt"
)

func eventualSafeNodes(graph [][]int) []int {
	var result []int
	colors := make([]int, len(graph))
	for i := range graph {
		if dfs(i, colors, graph) {
			result = append(result, i)
		}
	}
	return result
}

func dfs(start int, colors []int, graph [][]int) bool {
	if colors[start] != 0 {
		return colors[start] == 1
	}

	colors[start] = 2
	for _, node := range graph[start] {
		switch colors[node] {
		case 0:
			if !dfs(node, colors, graph) {
				return false
			}
		case 1:
			continue
		case 2:
			return false
		}
	}
	colors[start] = 1
	return true
}

func main() {
	graph := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	fmt.Println(eventualSafeNodes(graph))
}

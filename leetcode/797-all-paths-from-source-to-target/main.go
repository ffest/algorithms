package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	return findPath(&graph, 0)
}

func findPath(graph *[][]int, node int) [][]int {
	res := make([][]int, 0)
	if node == len(*graph)-1 {
		return [][]int{{node}}
	}
	for _, next := range (*graph)[node] {
		subRes := findPath(graph, next)
		for _, path := range subRes {
			res = append(res, append([]int{node}, path...))
		}
	}
	return res
}

func main() {
	var graph [][]int

	// example
	graph = append(graph, []int{1, 2})
	graph = append(graph, []int{3})
	graph = append(graph, []int{3})
	graph = append(graph, []int{})

	fmt.Println(allPathsSourceTarget(graph))
}

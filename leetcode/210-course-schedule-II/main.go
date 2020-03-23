package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	graph := make(map[int][]int)
	for _, preq := range prerequisites {
		want, need := preq[0], preq[1]
		graph[want] = append(graph[want], need)
	}
	visited := make(map[int]bool)
	using := make(map[int]bool)
	order := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		if hasCycles(i, graph, visited, using, &order) {
			return []int{}
		}
	}
	return order
}

func hasCycles(subj int, graph map[int][]int, visited, using map[int]bool, order *[]int) bool {
	if using[subj] {
		return true
	}
	if visited[subj] {
		return false
	}

	using[subj] = true
	for _, need := range graph[subj] {
		if hasCycles(need, graph, visited, using, order) {
			return true
		}
	}
	visited[subj] = true
	using[subj] = false
	*order = append(*order, subj)
	return false
}

func main() {
	numCourses := 4
	prerequisites := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}
	fmt.Println(findOrder(numCourses, prerequisites))
}

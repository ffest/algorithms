package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses < 2 || len(prerequisites) == 0 {
		return true
	}

	adjMatrix := make(map[int][]int)
	for _, prereq := range prerequisites {
		want := prereq[0]
		need := prereq[1]
		adjMatrix[want] = append(adjMatrix[want], need)
	}

	visited := make(map[int]bool)
	using := make(map[int]bool)

	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}

		if hasCycles(i, adjMatrix, visited, using) {
			return false
		}
	}

	return true
}

func hasCycles(subj int, adjMatrix map[int][]int, visited, using map[int]bool) bool {
	if using[subj] {
		return true
	}

	needs, ok := adjMatrix[subj]
	if !ok {
		visited[subj] = true
		return false
	}

	for _, need := range needs {
		using[subj] = true
		if hasCycles(need, adjMatrix, visited, using) {
			return true
		}
	}

	visited[subj] = true
	using[subj] = false

	return false
}

func main() {
	numCourses := 3
	prerequisites := [][]int{
		{1, 0},
		{2, 0},
	}
	fmt.Println(canFinish(numCourses, prerequisites))
}

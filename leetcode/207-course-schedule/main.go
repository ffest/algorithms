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

	using := make(map[int]bool)
	visited := make(map[int]bool)

	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		if hasCycles(i, adjMatrix, using, visited) {
			return false
		}
	}

	return true
}

func hasCycles(curSubject int, adjMatrix map[int][]int, using, visited map[int]bool) bool {
	if using[curSubject] {
		return true
	}

	needs, ok := adjMatrix[curSubject]
	if !ok {
		visited[curSubject] = true
		return false
	}

	for _, need := range needs {
		using[curSubject] = true
		if hasCycles(need, adjMatrix, using, visited) {
			return true
		}
	}
	using[curSubject] = false
	visited[curSubject] = true

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

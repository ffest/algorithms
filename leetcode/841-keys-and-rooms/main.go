package main

import (
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visited := make(map[int]struct{})

	queue := make([]int, 0)
	queue = append(queue, rooms[0]...)
	visited[0] = struct{}{}
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if _, ok := visited[next]; ok {
			continue
			visited[next] = struct{}{}
		}
		visited[next] = struct{}{}
		queue = append(queue, rooms[next]...)
	}

	return len(visited) == len(rooms)
}

func main() {
	//rooms := [][]int{{1}, {2}, {3}, {}}
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}

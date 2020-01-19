package main

import (
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visited := make(map[int]struct{})
	visited[0] = struct{}{}
	queue := append([]int{}, rooms[0]...)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if _, ok := visited[next]; ok {
			continue
		}
		visited[next] = struct{}{}
		queue = append(queue, rooms[next]...)
	}
	return len(rooms) == len(visited)
}

func main() {
	//rooms := [][]int{{1}, {2}, {3}, {}}
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}

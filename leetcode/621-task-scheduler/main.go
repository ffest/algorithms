package main

import (
	"fmt"
)

func leastInterval(tasks []byte, n int) int {
	cache := make(map[byte]int)

	var maxTask, count int
	for _, task := range tasks {
		cache[task]++
		if cache[task] > maxTask {
			count = 1
			maxTask = cache[task]
		} else if cache[task] == maxTask {
			count++
		}
	}

	idles := maxTask - 1
	idleSize := n - count + 1
	maxTasksSpace := maxTask * count
	return max(len(tasks), maxTasksSpace+idles*idleSize)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tasks := []byte{'A', 'A', 'B', 'B', 'A', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}

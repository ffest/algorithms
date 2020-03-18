package main

import (
	"fmt"
)

func prisonAfterNDays(cells []int, N int) []int {
	if len(cells) == 0 || N == 0 {
		return cells
	}

	foundCycle, cycleLen := false, 0
	cache := make(map[string]struct{})
	for N > 0 {
		N--
		cells = nextDay(cells)
		var key string
		for i := range cells {
			key += fmt.Sprintf("%d", cells[i])
		}
		if _, ok := cache[key]; ok {
			foundCycle = true
			break
		}
		cycleLen++
		cache[key] = struct{}{}
	}

	if foundCycle {
		N %= cycleLen
		for i := 0; i < N; i++ {
			cells = nextDay(cells)
		}
	}
	return cells
}

func nextDay(cells []int) []int {
	next := make([]int, len(cells))
	for i := 1; i < len(cells)-1; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		} else {
			next[i] = 0
		}
	}
	return next
}

func main() {
	cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	fmt.Println(prisonAfterNDays(cells, 7))
}

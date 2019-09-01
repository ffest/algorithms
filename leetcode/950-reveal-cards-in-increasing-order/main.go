package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	queue := make([]int, 0)
	n := len(deck)
	for i := 0; i < n; i++ {
		queue = append(queue, i)
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[queue[0]] = deck[i]
		queue = queue[1:]
		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}
	return result
}

func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	fmt.Println(deckRevealedIncreasing(deck))
}

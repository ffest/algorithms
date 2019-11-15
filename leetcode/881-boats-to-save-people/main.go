package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var count int
	i, j := 0, len(people)-1
	for ; i <= j; j-- {
		count++
		if people[i]+people[j] <= limit {
			i++
		}
	}
	return count
}

func main() {
	people := []int{3, 2, 2, 1}
	limit := 3
	fmt.Println(numRescueBoats(people, limit))
}

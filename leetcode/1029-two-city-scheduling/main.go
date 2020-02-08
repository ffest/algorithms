package main

import (
	"fmt"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	var sum int
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			sum += costs[i][0]
		} else {
			sum += costs[i][1]
		}
	}
	return sum
}

func main() {
	costs := [][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}
	fmt.Println(twoCitySchedCost(costs))
}

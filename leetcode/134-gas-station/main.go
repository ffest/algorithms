package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	var minIdx, sum int
	min := 1<<31 - 1
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < min {
			min = sum
			minIdx = i
		}
	}
	if sum < 0 {
		return -1
	}
	return (minIdx + 1) % len(gas)
}

func main() {
	gas := []int{3, 1, 1}
	cost := []int{1, 2, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

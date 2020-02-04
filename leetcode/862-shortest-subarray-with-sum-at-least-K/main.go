package main

import (
	"fmt"
)

func shortestSubarray(A []int, K int) int {
	sums := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		sums[i+1] = sums[i] + A[i]
	}
	min, dequeue := len(A)+1, make([]int, 0)
	for i := 0; i <= len(A); i++ {
		for len(dequeue) > 0 && sums[i]-sums[dequeue[0]] >= K {
			if min > i-dequeue[0] {
				min = i - dequeue[0]
			}
			dequeue = dequeue[1:]
		}
		for len(dequeue) > 0 && sums[dequeue[len(dequeue)-1]] >= sums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)
	}
	if min == len(A)+1 {
		return -1
	}
	return min
}

func main() {
	A := []int{84, -37, 32, 40, 95}
	K := 167
	fmt.Println(shortestSubarray(A, K))
}

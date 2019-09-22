package main

import (
	"container/heap"
	"fmt"
)

type maxHeap [][]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func kClosest(points [][]int, K int) [][]int {
	mh := maxHeap{}
	heap.Init(&mh)

	for _, point := range points {
		heap.Push(&mh, point)
		if mh.Len() > K {
			heap.Remove(&mh, 0)
		}
	}

	result := make([][]int, 0, K)
	for _, point := range mh {
		result = append(result, point)
	}
	return result
}

func main() {
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	fmt.Println(kClosest(points, 2))
}

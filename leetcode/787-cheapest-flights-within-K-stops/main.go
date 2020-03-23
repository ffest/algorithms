package main

import (
	"container/heap"
	"fmt"
)

type Path struct {
	node  int
	price int
	k     int
}

type minHeap []Path

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Path))
}

func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	cache := make(map[int]map[int]int)
	for _, flight := range flights {
		if _, ok := cache[flight[0]]; !ok {
			cache[flight[0]] = make(map[int]int)
		}
		cache[flight[0]][flight[1]] = flight[2]
	}
	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, Path{price: 0, node: src, k: K + 1})

	for h.Len() > 0 {
		path := heap.Pop(h).(Path)
		if path.node == dst {
			return path.price
		}
		if path.k > 0 {
			for next, price := range cache[path.node] {
				heap.Push(h, Path{node: next, price: path.price + price, k: path.k - 1})
			}
		}
	}
	return -1
}

func main() {
	n := 3
	edges := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	src := 0
	dst := 2
	k := 1
	fmt.Println(findCheapestPrice(n, edges, src, dst, k))
}

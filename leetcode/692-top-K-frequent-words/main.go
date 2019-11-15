package main

import (
	"fmt"
)

// Priority queue solution
/*type wf struct {
	w string
	f int
}
type maxHeap []wf

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i].f == h[j].f {
		return h[i].w > h[j].w
	}
	return h[i].f < h[j].f
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(wf))
}

func topKFrequent(words []string, k int) []string {
	cache := make(map[string]int)
	for _, w := range words {
		cache[w]++
	}

	var mh maxHeap
	heap.Init(&mh)
	for w, f := range cache {
		heap.Push(&mh, wf{w, f})
		if mh.Len() > k {
			heap.Pop(&mh)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(&mh).(wf).w
	}

	return result
}
*/

// TODO: bucket solution
func topKFrequent(words []string, k int) []string {
	return nil
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 3
	fmt.Println(topKFrequent(words, k))
}

package main

import (
	"fmt"
)

type info struct {
	number int
	freq   int
}

/*type minHeap []info

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(info))
}

func (h *minHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	for _, n := range nums {
		cache[n]++
	}

	minHeap := minHeap{}
	// Push to priority queue
	heap.Init(&minHeap)

	for num, freq := range cache {
		if len(minHeap) < k {
			heap.Push(&minHeap, info{number: num, freq: freq})
		} else if minHeap[0].freq < freq {
			minHeap[0] = info{number: num, freq: freq}
			heap.Fix(&minHeap, 0)
		}
	}
	result := make([]int, 0, k)
	for _, info := range minHeap {
		result = append(result, info.number)
	}
	return result
}*/

// Bucket sort
func topKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	for _, n := range nums {
		cache[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range cache {
		bucket := buckets[freq]
		if bucket == nil {
			bucket = []int{num}
		} else {
			bucket = append(bucket, num)
		}
		buckets[freq] = bucket
	}

	result := make([]int, 0, k)
	for i := len(nums); i > 0; i-- {
		bucket := buckets[i]
		if bucket == nil {
			continue
		} else {
			for _, num := range bucket {
				if k == 0 {
					return result
				}
				result = append(result, num)
				k--
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}

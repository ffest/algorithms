package main

import (
	"fmt"
)

type heap struct {
	buf  []int
	size int
}

func (h *heap) pop() int {
	if len(h.buf) == 0 {
		return 0
	}
	val := h.buf[0]
	h.buf[0], h.buf[len(h.buf)-1] = h.buf[len(h.buf)-1], h.buf[0]
	h.buf = h.buf[:len(h.buf)-1]
	h.down(0)
	return val
}

func (h *heap) push(x int) {
	if len(h.buf) == 0 || x > h.buf[0] {
		h.buf[0] = x
	}
	h.down(0)
}

func (h *heap) down(idx int) {
	leftChild := 2*idx + 1
	rightChild := 2*idx + 2

	switch {
	case rightChild < len(h.buf):
		if h.buf[rightChild] < h.buf[idx] && h.buf[rightChild] < h.buf[leftChild] {
			h.buf[idx], h.buf[rightChild] = h.buf[rightChild], h.buf[idx]
			h.down(rightChild)
		} else if h.buf[leftChild] < h.buf[idx] {
			h.buf[idx], h.buf[leftChild] = h.buf[leftChild], h.buf[idx]
			h.down(leftChild)
		}
	case leftChild < len(h.buf) && h.buf[leftChild] < h.buf[idx]:
		h.buf[idx], h.buf[leftChild] = h.buf[leftChild], h.buf[idx]
		h.down(leftChild)
	}
}

func (h *heap) build(arr []int) {
	for i := 0; i < h.size; i++ {
		h.buf = append(h.buf, arr[i])
	}
	for i := len(h.buf)/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	for i := h.size; i < len(arr); i++ {
		h.push(arr[i])
	}
}

func findKthLargest(nums []int, k int) int {
	h := new(heap)
	h.size = k
	h.build(nums)
	return h.pop()
}

func main() {
	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 5
	fmt.Println(findKthLargest(input, k))
}

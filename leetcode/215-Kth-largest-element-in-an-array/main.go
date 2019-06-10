package main

import (
	"fmt"
)

type heap struct {
	buf []int
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

func (h *heap) down(idx int) {
	rightChild := 2*idx + 2
	if rightChild < len(h.buf) && h.buf[rightChild] > h.buf[idx] {
		h.buf[idx], h.buf[rightChild] = h.buf[rightChild], h.buf[idx]
		h.down(rightChild)
	}
	leftChild := 2*idx + 1
	if leftChild < len(h.buf) && h.buf[leftChild] > h.buf[idx] {
		h.buf[idx], h.buf[leftChild] = h.buf[leftChild], h.buf[idx]
		h.down(leftChild)
	}
}

func (h *heap) build() {
	for i := len(h.buf)/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func findKthLargest(nums []int, k int) int {
	h := &heap{nums}
	h.build()

	var val int
	for i := 0; i < k; i++ {
		val = h.pop()
	}
	return val
}

func main() {
	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(input, k))
}

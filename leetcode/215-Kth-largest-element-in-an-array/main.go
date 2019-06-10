package main

import (
	"fmt"
	"log"
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
	leftChild := 2*idx + 1
	rightChild := 2*idx + 2

	switch {
	case rightChild < len(h.buf):
		if h.buf[rightChild] > h.buf[idx] && h.buf[rightChild] > h.buf[leftChild] {
			h.buf[idx], h.buf[rightChild] = h.buf[rightChild], h.buf[idx]
			h.down(rightChild)
		} else if h.buf[leftChild] > h.buf[idx] {
			h.buf[idx], h.buf[leftChild] = h.buf[leftChild], h.buf[idx]
			h.down(leftChild)
		}
	case leftChild < len(h.buf) && h.buf[leftChild] > h.buf[idx]:
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

	log.Print(h.buf)
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

package main

import "log"

type MinStack struct {
	helper []int
	stack  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		helper: make([]int, 0),
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if len(ms.helper) == 0 || ms.GetMin() >= x {
		ms.helper = append(ms.helper, x)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	if ms.stack[len(ms.stack)-1] == ms.GetMin() {
		ms.helper = ms.helper[:len(ms.helper)-1]
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.helper[len(ms.helper)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	log.Print(minStack.GetMin())
	minStack.Pop()
	log.Print(minStack.GetMin())
	minStack.Pop()
	log.Print(minStack.GetMin())
	minStack.Pop()
	log.Print(minStack.GetMin())
}

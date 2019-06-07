package main

import "log"

type MinStack struct {
	helper   []int
	stack    []int
	elements map[int]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		helper:   make([]int, 0),
		elements: make(map[int]int),
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)

	if len(ms.helper) == 0 || ms.helper[len(ms.helper)-1] > x {
		ms.helper = append(ms.helper, x)
	}
	if ms.helper[len(ms.helper)-1] == x {
		ms.elements[x]++
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	toPop := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	if toPop == ms.helper[len(ms.helper)-1] {
		if ms.elements[toPop] == 1 {
			ms.helper = ms.helper[:len(ms.helper)-1]
		}
		ms.elements[toPop]--
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.helper) == 0 {
		return 0
	}
	return ms.helper[len(ms.helper)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	log.Print(minStack.GetMin())
	minStack.Pop()
	log.Print(minStack.Top())
	log.Print(minStack.GetMin())
}

package main

import (
	"fmt"
	"math"
)

type MaxStack struct {
	max   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{
		max:   math.MinInt32,
		stack: make([]int, 0),
	}
}

// O(1)
func (ms *MaxStack) Push(x int) {
	if x >= ms.max {
		ms.stack = append(ms.stack, ms.max)
		ms.max = x
	}
	ms.stack = append(ms.stack, x)
}

// O(1)
func (ms *MaxStack) Pop() int {
	x := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	return x
}

// O(1)
func (ms *MaxStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

// O(1)
func (ms *MaxStack) PeekMax() int {
	return ms.max
}

// O(n)
func (ms *MaxStack) PopMax() int {
	tmp := make([]int, 0)
	for ms.Top() != ms.max {
		tmp = append(tmp, ms.Pop())
	}
	x := ms.Pop()
	ms.max = ms.Pop()
	for len(tmp) > 0 {
		val := tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		if val >= ms.max {
			ms.stack = append(ms.stack, ms.max)
			ms.max = val
		}
		ms.stack = append(ms.stack, val)
	}
	return x
}

func main() {
	ms := Constructor()
	ms.Push(5)
	ms.Push(1)
	ms.Push(5)
	fmt.Println(ms.Top())
	fmt.Println(ms.PopMax())
	fmt.Println(ms.Top())
	fmt.Println(ms.PeekMax())
	fmt.Println(ms.Pop())
	fmt.Println(ms.Top())
}

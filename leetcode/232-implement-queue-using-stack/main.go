package main

import "fmt"

type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	q.Peek()
	element := q.output[len(q.output)-1]
	q.output = q.output[:len(q.output)-1]
	return element
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.output) == 0 {
		for len(q.input) > 0 {
			q.output = append(q.output, q.input[len(q.input)-1])
			q.input = q.input[:len(q.input)-1]
		}
	}
	return q.output[len(q.output)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

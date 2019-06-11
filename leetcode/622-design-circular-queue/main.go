package main

import (
	"fmt"
)

type MyCircularQueue struct {
	queue  []int
	head   int
	tail   int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		tail:  -1,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		return false
	}
	cq.tail++
	cq.tail = cq.tail % len(cq.queue)
	cq.queue[cq.tail] = value
	cq.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}
	cq.head++
	cq.head = cq.head % len(cq.queue)
	cq.length--

	return true
}

/** Get the front item from the queue. */
func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.queue[cq.head]
}

/** Get the last item from the queue. */
func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.queue[cq.tail]
}

/** Checks whether the circular queue is empty or not. */
func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.length == 0
}

/** Checks whether the circular queue is full or not. */
func (cq *MyCircularQueue) IsFull() bool {
	return cq.length == len(cq.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	queue := Constructor(6)
	fmt.Println(queue.EnQueue(6))
	fmt.Println(queue.Rear())
	fmt.Println(queue.IsFull())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.Rear())
}

package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{0, head}

	slow, fast := start, start
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return start.Next
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 10,
					},
				},
			},
		},
	}
	fmt.Println(removeNthFromEnd(list, 3))
}

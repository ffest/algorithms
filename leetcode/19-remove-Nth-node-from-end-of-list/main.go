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
	start := &ListNode{0, nil}
	b, f := start, start
	b.Next = head
	for i := 0; i <= n; i++ {
		f = f.Next
	}

	for f != nil {
		b = b.Next
		f = f.Next
	}
	b.Next = b.Next.Next

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

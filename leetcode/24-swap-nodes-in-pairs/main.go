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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{0, head}
	current := result
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := first.Next

		tmp := second.Next
		second.Next = first
		first.Next = tmp
		current.Next = second
		current = current.Next.Next
	}

	return result.Next
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Print(swapPairs(list))
}

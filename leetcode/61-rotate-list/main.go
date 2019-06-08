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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	// Finding out the length of the list
	var length = 1
	for p.Next != nil {
		length++
		p = p.Next
	}
	p.Next = head

	for i := length - k%length; i > 1; i-- {
		head = head.Next
	}

	p = head.Next
	head.Next = nil

	return p
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
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	fmt.Println(rotateRight(list, 2))
}

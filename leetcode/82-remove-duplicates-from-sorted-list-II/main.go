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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, nil}
	tail := dummy

	var count = 1
	for head.Next != nil {
		if head.Val != head.Next.Val {
			if count == 1 {
				tail.Next = head
				tail = tail.Next
			} else {
				count = 1
			}
		} else {
			count++
		}
		head = head.Next
	}
	if count == 1 {
		tail.Next = head
		tail = tail.Next
	}
	tail.Next = nil
	return dummy.Next
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
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
			},
		},
	}

	fmt.Println(deleteDuplicates(list))
}

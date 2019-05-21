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

func middleNode(head *ListNode) *ListNode {
	list := head
	var count int
	for list != nil {
		count++
		list = list.Next
	}

	var iter int
	for {
		if iter == count/2 {
			return head
		}
		iter++
		head = head.Next
	}

	return nil
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
				},
			},
		},
	}
	fmt.Println(middleNode(list))
}

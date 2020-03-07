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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	node := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Println(mergeTwoLists(list1, list2))
}

package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less, greater := &ListNode{0, nil}, &ListNode{0, nil}
	p, q := less, greater

	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
	}
	q.Next = nil
	p.Next = greater.Next
	return less.Next
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	target := 3
	fmt.Println(partition(input, target))
}

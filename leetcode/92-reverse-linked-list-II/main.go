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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	result := &ListNode{0, head}

	left := result
	for i := 0; i < m-1; i++ {
		left = left.Next
	}

	middle := left.Next
	right := middle.Next

	for i := 0; i < n-m; i++ {
		middle.Next = right.Next
		right.Next = left.Next
		left.Next = right
		right = middle.Next
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
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(reverseBetween(list, 2, 4))
}

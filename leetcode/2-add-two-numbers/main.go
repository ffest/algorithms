package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	tail := &ListNode{0, nil}
	head := tail
	var num int
	for a != nil || b != nil {
		num /= 10
		if a != nil {
			num += a.Val
			a = a.Next
		}
		if b != nil {
			num += b.Val
			b = b.Next
		}

		head.Next = &ListNode{Val: num % 10}
		head = head.Next
	}

	if num/10 == 1 {
		head.Next = &ListNode{Val: 1}
	}

	return tail.Next
}

func main() {
	first := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	second := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Println(addTwoNumbers(first, second))
}

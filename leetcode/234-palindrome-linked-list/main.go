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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	for head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	/*list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 0,
		},
	}*/
	fmt.Println(isPalindrome(list))
}

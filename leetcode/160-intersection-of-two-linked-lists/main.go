package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != nil || q != nil {
		if p == q {
			return p
		}

		if p == nil {
			p = headB
			q = q.Next
			continue
		}

		if q == nil {
			q = headA
			p = p.Next
			continue
		}

		p = p.Next
		q = q.Next
	}
	return nil
}
func main() {
	commonHead := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: commonHead,
		},
	}

	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: commonHead,
			},
		},
	}
	fmt.Println(getIntersectionNode(headA, headB))
}

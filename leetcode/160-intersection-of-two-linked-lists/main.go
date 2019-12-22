package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Without calculating lengths of lists
/*func getIntersectionNode(headA, headB *ListNode) *ListNode {
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
}*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := lengthOfList(headA), lengthOfList(headB)
	if lenA < lenB {
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}
	for i := 0; i < lenA-lenB; i++ {
		headA = headA.Next
	}

	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func lengthOfList(l *ListNode) int {
	var length int
	for l != nil {
		l = l.Next
		length++
	}
	return length
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

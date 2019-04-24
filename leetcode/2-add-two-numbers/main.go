package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// A Nil Node simplifies additions in the main loop.
var empty = &ListNode{Val: 0}

// If the next node is nil in either number, use empty.
func nextOrEmpty(l *ListNode) *ListNode {
	if l.Next == nil {
		return empty
	}
	return l.Next
}

// Add two lists of digits together, each list having their least-significant digit in the same place.
func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	x := new(ListNode)
	head := x

	carry := 0
	for {
		x.Val = a.Val + b.Val + carry
		if x.Val > 9 {
			x.Val -= 10
			carry = 1
		} else {
			carry = 0
		}

		a, b = nextOrEmpty(a), nextOrEmpty(b)
		if a == empty && b == empty {
			break
		}

		x.Next = new(ListNode)
		x = x.Next
	}

	if carry > 0 {
		x.Next = &ListNode{Val: carry}
	}

	return head
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
				Val: 7,
			},
		},
	}
	fmt.Println(addTwoNumbers(first, second))
}

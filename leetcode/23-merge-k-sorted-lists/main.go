package main

import (
	"fmt"
)

/*
 Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := new(ListNode)
	subres := result
	for {
		min := 1<<31 - 1
		minIdx := 0
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if min > lists[i].Val {
				min = lists[i].Val
				minIdx = i
			}
		}
		next := lists[minIdx]
		if next == nil {
			break
		}
		subres.Next = next
		subres = subres.Next
		lists[minIdx] = lists[minIdx].Next
	}

	return result.Next
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
	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	lists := make([]*ListNode, 0)
	lists = append(lists, list1, list2, list3)
	fmt.Println(mergeKLists(lists))
}

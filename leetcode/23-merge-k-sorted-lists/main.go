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
	elements := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			elements = append(elements, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}
	if len(elements) == 0 {
		return nil
	}
	elements = Quick(elements)

	result := &ListNode{Val: elements[0]}
	subres := result

	for i := 1; i < len(elements); i++ {
		subres.Next = &ListNode{Val: elements[i]}
		subres = subres.Next
	}

	return result
}

func Quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := 0, len(arr)-1
	median := (1<<31 - 1) % len(arr)

	arr[median], arr[right] = arr[right], arr[median]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Quick(arr[:left])
	Quick(arr[left+1:])

	return arr
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

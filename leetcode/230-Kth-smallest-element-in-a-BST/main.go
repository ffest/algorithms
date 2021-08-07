package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }


func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return current.Val
		}
		current = current.Right
	}
	return 0
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(kthSmallest(root, 3))
}

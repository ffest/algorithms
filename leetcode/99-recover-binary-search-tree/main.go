package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	traverse(root, &first, &second, &prev)
	first.Val, second.Val = second.Val, first.Val
}

func traverse(root *TreeNode, first, second, prev **TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left, first, second, prev)
	if *prev != nil && (*prev).Val >= root.Val {
		if *first == nil {
			*first = *prev
		}
		*second = root
	}
	*prev = root
	traverse(root.Right, first, second, prev)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	recoverTree(root)
	fmt.Println(root)
}

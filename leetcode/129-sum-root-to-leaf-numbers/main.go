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

func sumNumbers(root *TreeNode) int {
	return helper(0, root)
}

func helper(current int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	current = 10*current + root.Val
	if root.Left == nil && root.Right == nil {
		return current
	}
	return helper(current, root.Left) + helper(current, root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(sumNumbers(root))
}

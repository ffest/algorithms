package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeftUnival := isUnivalTree(root.Left)
	isRightUnival := isUnivalTree(root.Right)

	if !isLeftUnival || !isRightUnival ||
		(root.Left != nil && root.Left.Val != root.Val) ||
		(root.Right != nil && root.Right.Val != root.Val) {
		return false
	}
	return true
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(isUnivalTree(root))
}

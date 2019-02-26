package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Test  string
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 0, Test: "left2"},
			Right: &TreeNode{Val: 0, Test: "left3"},
			Test:  "left1",
		},
		Right: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 0, Test: "right2"},
			Right: &TreeNode{Val: 1, Test: "right3"},
			Test:  "right1",
		},
		Test: "root",
	}
	fmt.Println(pruneTree(tree))
}

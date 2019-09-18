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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int
	height := fullHeight(root)
	for root != nil {
		height--
		rightHeight := fullHeight(root.Right)
		if height == rightHeight {
			count += 1 << uint(height)
			root = root.Right
		} else {
			count += 1 << uint(height-1)
			root = root.Left
		}
	}
	return count
}

func fullHeight(root *TreeNode) int {
	var height int
	for root != nil {
		height++
		root = root.Left
	}
	return height
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(countNodes(root))
}

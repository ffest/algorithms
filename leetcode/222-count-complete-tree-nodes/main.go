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
	fullHeight := findHeight(root)
	for root != nil {
		fullHeight--
		rightHeight := findHeight(root.Right)
		if rightHeight == fullHeight {
			count += 1 << uint(fullHeight)
			root = root.Right
		} else {
			count += 1 << uint(fullHeight-1)
			root = root.Left
		}
	}

	return count
}

func findHeight(root *TreeNode) int {
	var height int
	for root != nil {
		root = root.Left
		height++
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

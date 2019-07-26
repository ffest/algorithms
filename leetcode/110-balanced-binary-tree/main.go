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

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	fmt.Println(isBalanced(tree))
}

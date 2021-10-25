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

// bfs
/*func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		depth++
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left == nil && current.Right == nil {
				return depth
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return 0
}*/

// dfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		left = minDepth(root.Left)
	}
	if root.Right != nil {
		right = minDepth(root.Right)
	}
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1
}

func min(a, b int) int {
	if a < b {
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
	fmt.Println(minDepth(tree))
}

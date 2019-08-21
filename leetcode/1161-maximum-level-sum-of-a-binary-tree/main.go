package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 1

	maxSum := -math.MaxInt64
	maxSumLevel := 1
	for len(queue) > 0 {
		qSize := len(queue)
		current := 0
		for i := 0; i < qSize; i++ {
			node := queue[0]
			queue = queue[1:]
			current += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if current > maxSum {
			maxSum = current
			maxSumLevel = level
		}
		level++
	}

	return maxSumLevel
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: -8,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(maxLevelSum(root))
}

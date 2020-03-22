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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	traversal := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var level int
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if len(traversal) == level {
				traversal = append(traversal, []int{node.Val})
			} else {
				traversal[level] = append(traversal[level], node.Val)
			}
		}
		level++
	}
	return traversal
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
	fmt.Println(levelOrder(tree))
}

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

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	paths := make([][]int, 0)
	backtracking(root, []int{root.Val}, root.Val, sum, &paths)
	return paths
}

func backtracking(root *TreeNode, path []int, sum, target int, paths *[][]int) {
	if sum == target && root.Left == root.Right {
		*paths = append(*paths, path)
		return
	}
	if root.Left != nil {
		backtracking(root.Left, append(append([]int{}, path...), root.Left.Val), sum+root.Left.Val, target, paths)
	}
	if root.Right != nil {
		backtracking(root.Right, append(append([]int{}, path...), root.Right.Val), sum+root.Right.Val, target, paths)
	}
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}

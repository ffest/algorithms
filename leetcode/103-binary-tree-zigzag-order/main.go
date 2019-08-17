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
}

//TODO: level order with swapping
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	result = travel(root, result, 0)
	return result
}

func travel(tree *TreeNode, result [][]int, level int) [][]int {
	if level >= len(result) {
		subRes := make([]int, 0)
		subRes = append(subRes, tree.Val)
		result = append(result, subRes)
	} else {
		subRes := result[level]
		subRes = append(subRes, tree.Val)
		result[level] = subRes
	}
	level++

	if level%2 == 0 {
		if tree.Left != nil {
			result = travel(tree.Left, result, level)
			level--
		}
		if tree.Right != nil {
			result = travel(tree.Right, result, level)
			level--
		}
	} else {
		if tree.Right != nil {
			result = travel(tree.Right, result, level)
			level--
		}
		if tree.Left != nil {
			result = travel(tree.Left, result, level)
			level--
		}
	}

	return result
}

func main() {
	treeNode := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(treeNode))
}

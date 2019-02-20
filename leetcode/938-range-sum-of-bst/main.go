package main

import "fmt"

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

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	if root != nil {
		if root.Val >= L && root.Val <= R {
			sum += root.Val
		}
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

func main() {
	//example
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 15, Left: nil, Right: &TreeNode{Val: 18, Left: nil, Right: nil}},
	}
	fmt.Println(rangeSumBST(tree, 7, 15))
}

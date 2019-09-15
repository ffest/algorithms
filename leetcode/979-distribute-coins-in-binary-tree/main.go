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

// Recursive solution
/*func distributeCoins(root *TreeNode) int {
	var moves int
	if root.Left != nil {
		moves += distributeCoins(root.Left)
		root.Val += root.Left.Val - 1
		moves += abs(root.Left.Val - 1)
	}
	if root.Right != nil {
		moves += distributeCoins(root.Right)
		root.Val += root.Right.Val - 1
		moves += abs(root.Right.Val - 1)
	}
	return moves
}*/

// Recursive solution with helper
func distributeCoins(root *TreeNode) int {
	var moves int
	traverse(root, &moves)
	return moves
}

func traverse(root *TreeNode, moves *int) int {
	if root == nil {
		return 0
	}
	left, right := traverse(root.Left, moves), traverse(root.Right, moves)
	*moves += abs(left) + abs(right)
	return root.Val - 1 + left + right
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(distributeCoins(root))
}

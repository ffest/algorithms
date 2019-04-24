package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1
	return validate(min, max, root)
}

func validate(min, max int, node *TreeNode) bool {
	if node == nil {
		return true
	}

	return min < node.Val && max > node.Val &&
		validate(min, node.Val, node.Left) && validate(node.Val, max, node.Right)
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isValidBST(root))
}

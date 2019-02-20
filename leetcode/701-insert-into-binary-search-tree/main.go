package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func main() {
	root := TreeNode{Val: 5}
	insertIntoBST(&root, 10)
	insertIntoBST(&root, 2)
	insertIntoBST(&root, 9)

	log.Printf("%+v", root.Right)
}

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

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isEqual(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == t {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}

func main() {
	mini := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	tree := &TreeNode{
		Val:  3,
		Left: mini,
		Right: &TreeNode{
			Val: 5,
		},
	}

	fmt.Println(isSubtree(tree, mini))
}

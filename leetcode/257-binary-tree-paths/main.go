package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := make([]string, 0)
	helper(root, strconv.Itoa(root.Val), &paths)
	return paths
}

func helper(root *TreeNode, path string, paths *[]string) {
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
	}
	if root.Right != nil {
		helper(root.Right, path+"->"+strconv.Itoa(root.Right.Val), paths)
	}
	if root.Left != nil {
		helper(root.Left, path+"->"+strconv.Itoa(root.Left.Val), paths)
	}
}

func main() {
	treeNode := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 3,
				},
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
	fmt.Println(binaryTreePaths(treeNode))
}

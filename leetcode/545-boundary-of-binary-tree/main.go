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

// Mix of preorder and postorder traversals
func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := make([]int, 0)
	nodes = append(nodes, root.Val)
	getBoundary(root.Left, true, false, &nodes)
	getBoundary(root.Right, false, true, &nodes)
	return nodes
}

func getBoundary(root *TreeNode, lb, rb bool, nodes *[]int) {
	if root == nil {
		return
	}
	if lb {
		*nodes = append(*nodes, root.Val)
	}
	if !lb && !rb && root.Left == nil && root.Right == nil {
		*nodes = append(*nodes, root.Val)
	}
	getBoundary(root.Left, lb, rb && root.Right == nil, nodes)
	getBoundary(root.Right, lb && root.Left == nil, rb, nodes)
	if rb {
		*nodes = append(*nodes, root.Val)
	}
}

// Simple solution of 3 obvious steps
/*func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []int{root.Val}
	addLeftBoundaries(root.Left, &nodes)
	addLeaves(root.Left, &nodes)
	addLeaves(root.Right, &nodes)
	addRightBoundaries(root.Right, &nodes)
	return nodes
}

func addLeftBoundaries(root *TreeNode, nodes *[]int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	*nodes = append(*nodes, root.Val)
	if root.Left != nil {
		addLeftBoundaries(root.Left, nodes)
	} else if root.Right != nil {
		addLeftBoundaries(root.Right, nodes)
	}
}

func addRightBoundaries(root *TreeNode, nodes *[]int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	if root.Right != nil {
		addRightBoundaries(root.Right, nodes)
	} else if root.Left != nil {
		addRightBoundaries(root.Left, nodes)
	}
	*nodes = append(*nodes, root.Val)
}

func addLeaves(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*nodes = append(*nodes, root.Val)
	}
	addLeaves(root.Left, nodes)
	addLeaves(root.Right, nodes)
}*/

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(boundaryOfBinaryTree(root))
}

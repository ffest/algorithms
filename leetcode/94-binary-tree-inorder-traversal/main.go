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
/*func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	traversal := make([]int, 0)
	dfs(root, &traversal)
	return traversal
}

func dfs(root *TreeNode, traversal *[]int) {
	if root.Left != nil {
		dfs(root.Left, traversal)
	}
	*traversal = append(*traversal, root.Val)
	if root.Right != nil {
		dfs(root.Right, traversal)
	}
}*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	traversal := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		traversal = append(traversal, current.Val)
		current = current.Right
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
	fmt.Println(inorderTraversal(tree))
}

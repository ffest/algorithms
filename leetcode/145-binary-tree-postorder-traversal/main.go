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

// iterative solution
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	traversal := make([]int, 0)
	stack := make([]*TreeNode, 0)

	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			traversal = append(traversal, current.Val)
			stack = append(stack, current)
			current = current.Right
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		current = current.Left
	}
	reverse(traversal)
	return traversal
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

// recursive solution
func postorderTraversalRecursive(root *TreeNode) []int {
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
	if root.Right != nil {
		dfs(root.Right, traversal)
	}
	*traversal = append(*traversal, root.Val)
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
	fmt.Println(postorderTraversal(tree))
}

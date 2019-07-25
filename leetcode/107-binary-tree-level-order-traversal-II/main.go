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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	traversal := make([][]int, 0)
	helper(root, 0, &traversal)
	return traversal
}

func helper(root *TreeNode, level int, traversal *[][]int) {
	if root == nil {
		return
	}
	if len(*traversal) < level+1 {
		*traversal = append([][]int{{}}, *traversal...)
	}
	helper(root.Left, level+1, traversal)
	helper(root.Right, level+1, traversal)
	(*traversal)[len(*traversal)-level-1] = append((*traversal)[len(*traversal)-level-1], root.Val)
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
	fmt.Println(levelOrderBottom(tree))
}

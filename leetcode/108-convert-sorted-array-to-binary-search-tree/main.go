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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	median := len(nums) / 2
	return &TreeNode{
		Val:   nums[median],
		Left:  sortedArrayToBST(nums[:median]),
		Right: sortedArrayToBST(nums[median+1:]),
	}
}

func BSTToSortedArray(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nums = append(nums, root.Val)
			root = root.Right
		}
	}
	return nums
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Println(BSTToSortedArray(tree))
}

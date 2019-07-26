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

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}

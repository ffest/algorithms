package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// Base case.
	if len(nums) < 1 {
		return nil
	}

	// Find the max num.
	maxIndex, maxValue := max(nums)

	// Construct the next level.
	return &TreeNode{
		Val:   maxValue,
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}

func max(nums []int) (int, int) {
	var maxValue, maxIndex int
	for k, v := range nums {
		if v > maxValue {
			maxValue = v
			maxIndex = k
		}
	}
	return maxIndex, maxValue
}

func main() {
	input := []int{3, 2, 1, 6, 0, 5}
	fmt.Println(constructMaximumBinaryTree(input))
}

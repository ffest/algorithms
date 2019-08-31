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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	cache := make(map[string]int)
	result := make([]*TreeNode, 0)
	helper(root, cache, &result)
	return result
}

func helper(root *TreeNode, cache map[string]int, result *[]*TreeNode) string {
	if root == nil {
		return "|"
	}
	left, right := helper(root.Left, cache, result), helper(root.Right, cache, result)
	key := strconv.Itoa(root.Val) + "|" + left + "|" + right
	if cache[key] == 1 {
		*result = append(*result, root)
	}
	cache[key]++
	return key
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(findDuplicateSubtrees(root))
}

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generator(1, n)
}

func generator(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	list := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftSubs := generator(start, i-1)
		rightSubs := generator(i+1, end)
		for _, left := range leftSubs {
			for _, right := range rightSubs {
				list = append(list, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return list
}

func main() {
	n := 3
	fmt.Println(generateTrees(n))
}

package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreePosition struct {
	node     *TreeNode
	position int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := make(map[int]map[int][]int)
	queue := make([]TreePosition, 0)
	queue = append(queue, TreePosition{node: root, position: 0})
	deep := 0
	for len(queue) > 0 {
		currentLen := len(queue)
		for i := 0; i < currentLen; i++ {
			tp := queue[0]
			queue = queue[1:]

			if tp.node.Left != nil {
				queue = append(queue, TreePosition{node: tp.node.Left, position: tp.position - 1})
			}
			if tp.node.Right != nil {
				queue = append(queue, TreePosition{node: tp.node.Right, position: tp.position + 1})
			}

			levelNodes, ok := levels[tp.position]
			if !ok {
				levelNodes = make(map[int][]int)
			}
			levelNodes[deep] = append(levelNodes[deep], tp.node.Val)
			levels[tp.position] = levelNodes
		}
		deep++
	}

	result := make([][]int, 0)
	for i := -1000; i <= 1000; i++ {
		level, ok := levels[i]
		if !ok {
			continue
		}
		levelVals := make([]int, 0)
		for j := 0; j <= deep; j++ {
			vals, ok := level[j]
			if !ok {
				continue
			}
			sort.Ints(vals)
			levelVals = append(levelVals, vals...)
		}
		result = append(result, levelVals)
	}
	return result
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
	fmt.Println(verticalTraversal(tree))
}

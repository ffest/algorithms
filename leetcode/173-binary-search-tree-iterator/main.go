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

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var bi BSTIterator
	bi.push(root)
	return bi
}

func (bi *BSTIterator) push(root *TreeNode) {
	for root != nil {
		bi.stack = append(bi.stack, root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (bi *BSTIterator) Next() int {
	tmp := bi.stack[len(bi.stack)-1]
	bi.stack = bi.stack[:len(bi.stack)-1]
	bi.push(tmp.Right)
	return tmp.Val
}

/** @return whether we have a next smallest number */
func (bi *BSTIterator) HasNext() bool {
	if len(bi.stack) == 0 {
		return false
	}
	return true
}

func main() {
	tree := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	iter := Constructor(tree)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
}

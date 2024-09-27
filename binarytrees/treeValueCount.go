package binarytrees

import (
	"bulentkeskin.com/dsa/library"
)

/*
Write a function that takes in the root of a binary tree and a target value.
The function should return the number of times that the target occurs in the tree.
*/

func TreeValueCount(root *TreeNode, val int) int {
	var stack library.Stack[*TreeNode]

	stack.Push(root)
	res := 0

	for stack.Len() > 0 {
		node := stack.Pop()

		if node.Val == val {
			res++
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	return res
}

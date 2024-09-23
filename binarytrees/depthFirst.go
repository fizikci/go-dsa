package binarytrees

import "bulentkeskin.com/dsa/library"

func DepthFirst(root *TreeNode) []int {
	var stack library.Stack[*TreeNode]
	stack.Push(root)
	var values []int

	for stack.Len() > 0 {
		node := stack.Pop()

		values = append(values, node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return values
}

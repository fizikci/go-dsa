package binarytrees

import "bulentkeskin.com/dsa/library"

/*
Write a function, bottomRightValue, that takes in the root of a binary tree.
The function should return the right-most value in the bottom-most level of the tree.

You may assume that the input tree is non-empty.
*/

func BottomRightValue(root *TreeNode) int {
	var queue library.Queue[*TreeNode]
	queue.Enqueue(root)

	var node *TreeNode

	for queue.Len() > 0 {
		node = queue.Dequeue()

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return node.Val
}

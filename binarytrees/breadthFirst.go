package binarytrees

import "bulentkeskin.com/dsa/library"

/*
Write a function that takes in the root of a binary tree.
The function should return an array containing all values of the tree in breadth-first order.
*/

func BreadthFirst(root *TreeNode) []int {
	var queue library.Queue[*TreeNode]
	queue.Enqueue(root)

	var res []int

	for queue.Len() > 0 {
		node := queue.Dequeue()

		res = append(res, node.Val)

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return res
}

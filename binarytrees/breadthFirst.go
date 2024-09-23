package binarytrees

import "bulentkeskin.com/dsa/library"

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

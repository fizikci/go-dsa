package binarytrees

import "math"

/*
Write a function that takes in the root of a binary tree that contains number values.
The function should return the maximum value within the tree.
*/

func MaxValue(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}

	return int(math.Max(float64(root.Val), math.Max(float64(MaxValue(root.Left)), float64(MaxValue(root.Right)))))
}

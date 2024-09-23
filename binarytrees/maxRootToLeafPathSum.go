package binarytrees

import "math"

/*
Write a function that takes in the root of a binary tree that contains number values.
The function should return the maximum sum of any root to leaf path within the tree.
*/

func MaxRootToLeafPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	leftSum := root.Val + MaxRootToLeafPathSum(root.Left)
	rightSum := root.Val + MaxRootToLeafPathSum(root.Right)

	return int(math.Max(float64(leftSum), float64(rightSum)))
}

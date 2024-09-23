package binarytrees

import "math"

func MaxValue(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}

	return int(math.Max(float64(root.Val), math.Max(float64(MaxValue(root.Left)), float64(MaxValue(root.Right)))))
}

package binarytrees

/*
Write a function that takes in the root of a binary tree. The function should return
a 2-Dimensional array where each subarray represents a root-to-leaf path in the tree.

The order within an individual path must start at the root and end at the leaf,
but the relative order among paths in the outer array does not matter.

You may assume that the input tree is non-empty.
*/

func AllTreePaths(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	left := AllTreePaths(root.Left)
	for i, arr := range left {
		left[i] = append([]int{root.Val}, arr...)
	}

	right := AllTreePaths(root.Right)
	for i, arr := range right {
		right[i] = append([]int{root.Val}, arr...)
	}

	return append(append([][]int{}, left...), right...)
}

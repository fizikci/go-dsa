package binarytrees

/*
Write a function that takes in the root of a binary tree and a target value.
The function should return an array representing a path to the target value.
If the target value is not found in the tree, then return null.
*/

func FindPath(root *TreeNode, target int) []int {
	if root == nil {
		return nil
	}

	if root.Val == target {
		return []int{root.Val}
	}

	leftPath := FindPath(root.Left, target)
	if len(leftPath) > 0 {
		return append([]int{root.Val}, leftPath...)
	}

	rightPath := FindPath(root.Right, target)
	if len(rightPath) > 0 {
		return append([]int{root.Val}, rightPath...)
	}

	return nil
}

package binarytrees

/*
Write a function that takes in the root of a binary tree that contains number values.
The function should return the total sum of all values in the tree.
*/

func TreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + TreeSum(root.Left) + TreeSum(root.Right)
}

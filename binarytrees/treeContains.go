package binarytrees

/*
Write a function that takes in the root of a binary tree and a target value.
The function should return a boolean indicating whether or not the value is contained in the tree.
*/

func TreeContains(root *TreeNode, n int) bool {
	if root == nil {
		return false
	}

	return root.Val == n || TreeContains(root.Left, n) || TreeContains(root.Right, n)
}

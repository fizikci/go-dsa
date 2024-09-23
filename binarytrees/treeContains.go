package binarytrees

func TreeContains(root *TreeNode, n int) bool {
	if root == nil {
		return false
	}

	return root.Val == n || TreeContains(root.Left, n) || TreeContains(root.Right, n)
}

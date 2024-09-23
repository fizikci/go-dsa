package binarytrees

func TreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + TreeSum(root.Left) + TreeSum(root.Right)
}

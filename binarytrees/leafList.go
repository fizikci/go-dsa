package binarytrees

func LeafList(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	return append(LeafList(root.Left), LeafList(root.Right)...)
}

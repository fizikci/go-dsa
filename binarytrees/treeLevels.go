package binarytrees

/*
Write a function that takes in the root of a binary tree. The function should return
a 2-Dimensional array where each subarray represents a level of the tree.
*/

func TreeLevels(root *TreeNode) [][]int {
	res := [][]int{}

	var traverseLevels func(*TreeNode, int)

	traverseLevels = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)

		traverseLevels(node.Left, level+1)
		traverseLevels(node.Right, level+1)
	}

	traverseLevels(root, 0)

	return res
}

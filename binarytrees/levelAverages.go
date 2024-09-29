package binarytrees

/*
Write a function, levelAverages, that takes in the root of a binary tree that contains number values.
The function should return an array containing the average value of each level.
*/

func LevelAverages(root *TreeNode) []float32 {
	levels := [][]int{}

	var findLevels func(*TreeNode, int)
	findLevels = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(levels) == level {
			levels = append(levels, []int{})
		}

		levels[level] = append(levels[level], node.Val)

		findLevels(node.Left, level+1)
		findLevels(node.Right, level+1)
	}

	findLevels(root, 0)

	res := []float32{}

	for _, level := range levels {
		res = append(res, float32(sum(level))/float32(len(level)))
	}

	return res
}

func sum(level []int) int {
	if len(level) == 0 {
		return 0
	}
	return level[0] + sum(level[1:])
}

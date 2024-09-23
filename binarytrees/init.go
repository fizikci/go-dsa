package binarytrees

import (
	"strconv"
	"strings"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// treeToString returns a string representation of the binary tree.
func treeToString(root *TreeNode) string {
	if root == nil {
		return ""
	}
	lines, _, _, _ := buildTreeString(root)
	return strings.Join(lines, "\n")
}

// buildTreeString builds the string representation of the tree recursively.
func buildTreeString(node *TreeNode) ([]string, int, int, int) {
	if node == nil {
		return []string{}, 0, 0, 0
	}

	// Convert node value to string
	nodeStr := strconv.Itoa(node.Val)
	nodeWidth := len(nodeStr)

	// Recursively get string representations of left and right subtrees
	leftLines, leftPos, leftWidth, leftHeight := buildTreeString(node.Left)
	rightLines, rightPos, rightWidth, rightHeight := buildTreeString(node.Right)

	// Prepare variables to combine left and right subtrees
	firstLine := ""
	secondLine := ""
	var gapSize int

	// Determine the positions and sizes
	if leftWidth > 0 {
		firstLine += strings.Repeat(" ", leftPos) + strings.Repeat(" ", leftWidth-leftPos)
		secondLine += strings.Repeat(" ", leftPos) + "/" + strings.Repeat(" ", leftWidth-leftPos-1)
	} else {
		gapSize = nodeWidth
	}

	firstLine += nodeStr
	secondLine += strings.Repeat(" ", nodeWidth)

	if rightWidth > 0 {
		firstLine += strings.Repeat(" ", rightPos) + strings.Repeat(" ", rightWidth-rightPos)
		secondLine += strings.Repeat(" ", rightPos) + "\\" + strings.Repeat(" ", rightWidth-rightPos-1)
	} else {
		gapSize += nodeWidth
	}

	// Merge left and right subtree lines
	mergedLines := []string{firstLine, secondLine}
	for i := 0; i < max(leftHeight, rightHeight); i++ {
		var leftLine, rightLine string
		if i < leftHeight {
			leftLine = leftLines[i]
		} else {
			leftLine = strings.Repeat(" ", leftWidth)
		}
		if i < rightHeight {
			rightLine = rightLines[i]
		} else {
			rightLine = strings.Repeat(" ", rightWidth)
		}
		mergedLines = append(mergedLines, leftLine+strings.Repeat(" ", nodeWidth)+rightLine)
	}

	return mergedLines, leftWidth + nodeWidth/2, leftWidth + nodeWidth + rightWidth, len(mergedLines)
}

// Helper function to get the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (root *TreeNode) String() (res string) {
	return treeToString(root)
}

func ArrayToBinaryTree(arr []int) *TreeNode {
	// Base case: if array is empty, return nil
	if len(arr) == 0 {
		return nil
	}

	// Helper function to recursively build the tree
	var buildTree func(int) *TreeNode
	buildTree = func(i int) *TreeNode {
		if i >= len(arr) || arr[i] == -1 {
			return nil
		}
		// Create a new node for the current index
		node := &TreeNode{
			Val: arr[i],
		}
		// Recursively build left and right subtrees
		node.Left = buildTree(2*i + 1)
		node.Right = buildTree(2*i + 2)
		return node
	}

	// Start building the tree from the root (index 0)
	return buildTree(0)
}

var (
	IntBinaryTree *TreeNode
)

func init() {
	IntBinaryTree = ArrayToBinaryTree([]int{1, 2, 3, 4, 5, -1, 7})
}

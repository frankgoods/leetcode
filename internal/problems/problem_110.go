package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	var balanced = true
	check(root, 0, &balanced)
	return balanced
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(node *TreeNode, height int, balanced *bool) int {
	if node == nil || !(*balanced) {
		return height
	}

	hl := check(node.Left, height+1, balanced)
	hr := check(node.Right, height+1, balanced)
	if abs(hl-hr) > 1 {
		*balanced = false
	}

	return max(hl, hr)
}

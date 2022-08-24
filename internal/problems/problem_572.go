package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSame(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return (root != nil) &&
		(isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

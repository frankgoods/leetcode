package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func createInt(i int) *int {
	return &i
}

func isValidBST(root *TreeNode) bool {
	return isValid(root.Left, nil, &root.Val, root.Val) &&
		isValid(root.Right, &root.Val, nil, root.Val)
}

func isValid(node *TreeNode, l, r *int, parentVal int) bool {
	if node == nil {
		return true
	}

	if l != nil && node.Val <= *l {
		return false
	}

	if r != nil && node.Val >= *r {
		return false
	}

	return isValid(node.Left, l, &node.Val, node.Val) &&
		isValid(node.Right, &node.Val, r, node.Val)
}

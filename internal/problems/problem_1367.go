package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Visited map[*TreeNode]bool

func isSubPath(head *ListNode, root *TreeNode) bool {
	var res bool
	visited := make(Visited)
	subPath(head, root, head, visited, &res)
	return res
}

func subPath(listNode *ListNode, treeNode *TreeNode, listHead *ListNode, visited Visited, res *bool) {
	if *res {
		return
	}
	if listNode == nil {
		*res = true
		return
	}

	if treeNode == nil {
		return
	}

	if listNode.Val == treeNode.Val {
		subPath(listNode.Next, treeNode.Left, listHead, visited, res)
		subPath(listNode.Next, treeNode.Right, listHead, visited, res)
	}

	if _, ok := visited[treeNode]; ok {
		return
	} else {
		visited[treeNode] = true
		subPath(listHead, treeNode.Left, listHead, visited, res)
		subPath(listHead, treeNode.Right, listHead, visited, res)
	}
}

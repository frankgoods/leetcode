package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{-1, head}
	nthNode := preHead
	curNode := head
	curIdx := 1

	for curNode.Next != nil {
		curNode = curNode.Next
		curIdx++
		if curIdx > n {
			nthNode = nthNode.Next
		}
	}

	nthNode.Next = nthNode.Next.Next
	return preHead.Next
}

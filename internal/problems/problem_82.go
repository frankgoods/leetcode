package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Val: 0, Next: nil}

	curNew := newHead
	curOld := head

	for curOld != nil && curOld.Next != nil {
		if curOld.Val == curOld.Next.Val {
			val := curOld.Val
			for curOld != nil && curOld.Val == val {
				curOld = curOld.Next
			}
		} else {
			curNew.Next = curOld
			curOld = curOld.Next
			curNew = curNew.Next
		}
	}

	curNew.Next = curOld
	return newHead.Next
}

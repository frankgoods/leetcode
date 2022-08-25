package problems

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	cur := head
	for cur != nil {
		if _, exists := m[cur]; exists {
			return cur
		} else {
			m[cur] = true
			cur = cur.Next
		}
	}

	return nil
}

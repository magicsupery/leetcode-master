package list

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next

	}

	return prev
}

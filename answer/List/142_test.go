package list

func detectCycleUseMap(head *ListNode) *ListNode {
	checkMap := make(map[*ListNode]bool)

	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := checkMap[cur]; ok {
			return cur
		}

		checkMap[cur] = true
	}

	return nil
}

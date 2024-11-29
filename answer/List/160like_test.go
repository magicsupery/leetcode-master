package list

func getIntersectionNodeUseMap(headA, headB *ListNode) *ListNode {
	commonNodeMap := make(map[*ListNode]bool)

	for cur := headA; cur != nil; cur = cur.Next {
		commonNodeMap[cur] = true
	}

	for cur := headB; cur != nil; cur = cur.Next {
		if _, ok := commonNodeMap[cur]; ok {
			return cur
		}
	}

	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	for cur := headA; cur != nil; cur = cur.Next {
		lenA++
	}

	for cur := headB; cur != nil; cur = cur.Next {
		lenB++
	}

	curA, curB := headA, headB
	if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			curB = curB.Next
		}

	} else if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			curA = curA.Next
		}
	}

	for curA != nil {
		if curA == curB {
			return curA
		}

		curA = curA.Next
		curB = curB.Next
	}

	return nil

}

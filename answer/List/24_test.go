package list

import "testing"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pivot := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev := pivot
	var cur1, cur2 *ListNode
	var next *ListNode
	for prev != nil {
		cur1 = prev.Next
		if cur1 != nil {
			cur2 = cur1.Next
		} else {
			break
		}

		if cur2 != nil {
			next = cur2.Next
		} else {
			break
		}

		prev.Next = cur2
		cur2.Next = cur1
		cur1.Next = next

		prev = cur1
	}

	newHead := pivot.Next
	pivot.Next = nil
	return newHead
}

func Test_24(t *testing.T) {
	head := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	if head.Val != 2 || head.Next.Val != 1 || head.Next.Next.Val != 4 || head.Next.Next.Next.Val != 3 {
		t.Fail()
	}

	head2 := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	if head2.Val != 2 || head2.Next.Val != 1 || head2.Next.Next.Val != 3 {
		t.Fail()
	}

	head3 := swapPairs(&ListNode{1, nil})
	if head3.Val != 1 {
		t.Fail()
	}

	head4 := swapPairs(nil)
	if head4 != nil {
		t.Fail()
	}
}

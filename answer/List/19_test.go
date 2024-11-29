package list

import "testing"

func removeNthFromEndTwoPass(head *ListNode, n int) *ListNode {
	pivot := &ListNode{
		Val:  -1,
		Next: head,
	}

	length := 0
	for cur := pivot.Next; cur != nil; cur = cur.Next {
		length++
	}

	index := length - n

	var cur *ListNode
	prev := pivot
	cur = prev.Next
	for i := 1; i <= index; i++ {
		prev = cur
		cur = prev.Next
	}

	prev.Next = cur.Next
	cur.Next = nil

	return pivot.Next

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pivot := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev := pivot
	first := prev.Next
	second := first.Next
	for i := 1; i < n; i++ {
		second = second.Next
	}

	for second != nil {
		second = second.Next
		prev = prev.Next
		first = first.Next
	}

	prev.Next = first.Next
	first.Next = nil

	return pivot.Next
}

func Test_19(t *testing.T) {
	res1 := removeNthFromEndTwoPass(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)
	if res1.Val != 1 || res1.Next.Val != 2 || res1.Next.Next.Val != 3 || res1.Next.Next.Next.Val != 5 {
		t.Errorf("res1: %v", res1)
	}

	res2 := removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)
	if res2.Val != 1 || res2.Next.Val != 2 || res2.Next.Next.Val != 3 || res2.Next.Next.Next.Val != 5 {
		t.Errorf("res2: %v", res2)
	}
}

package list

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	for p := dummy; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

func Test_removeElements(t *testing.T) {
	tests := []struct {
		head   *ListNode
		val    int
		result *ListNode
	}{
		{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, 2, &ListNode{1, &ListNode{1, nil}}},
	}

	for _, test := range tests {
		result := removeElements(test.head, test.val)
		if !reflect.DeepEqual(result, test.result) {
			panic(test)
		}
	}
}

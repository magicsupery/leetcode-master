package list

import "testing"
import "github.com/stretchr/testify/assert"

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  -1,
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	for cur := this.Next; cur != nil && index >= 0; {
		if index == 0 {
			return cur.Val
		}
		cur = cur.Next
		index--
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	prev := this.Next
	this.Next = &MyLinkedList{
		Val:  val,
		Next: prev,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := &MyLinkedList{
		Val:  val,
		Next: nil,
	}

	for cur := this; cur != nil; {
		if cur.Next == nil {
			cur.Next = tail
			break
		}
		cur = cur.Next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev := this
	cur := prev.Next
	for cur != nil {
		if index == 0 {
			node := &MyLinkedList{
				Val:  val,
				Next: cur,
			}
			prev.Next = node
			break
		}

		prev = cur
		cur = cur.Next
		index--
	}
	//index == list length
	if index == 0 && cur == nil {
		prev.Next = &MyLinkedList{
			Val:  val,
			Next: nil,
		}
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev := this
	cur := prev.Next
	for cur != nil {
		if index == 0 {
			prev.Next = cur.Next
			break
		}

		prev = cur
		cur = cur.Next
		index--
	}
}

func Test_basic(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)

	assert.Equal(t, 1, obj.Get(0))
	assert.Equal(t, 2, obj.Get(1))
	assert.Equal(t, 3, obj.Get(2))
	obj.DeleteAtIndex(1)
	assert.Equal(t, 3, obj.Get(1))
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

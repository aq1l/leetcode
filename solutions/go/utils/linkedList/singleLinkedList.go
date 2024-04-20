package linkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListFromSlice takes int slice array and returns head node for singly linked list
func NewListFromSlice(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{}
	head.Val = data[0]
	prev := head

	data = data[1:]
	for i := range data {
		next := &ListNode{}
		next.Val = data[i]
		prev.Next = next
		prev = next
	}

	return head
}

func (l *ListNode) String() string {
	s := "["
	tempHead := l
	for tempHead != nil {
		s += fmt.Sprintf("%v", tempHead.Val)

		if tempHead.Next != nil {
			s += ", "
		}

		tempHead = tempHead.Next
	}

	return s + "]"
}

// CompareLists - returns boolean which indicates if all nodes' value is equal in passed two lists.
// Doesn't edit the passed lists.
func CompareLists(l1 *ListNode, l2 *ListNode) bool {
	tempL1, tempL2 := l1, l2
	for tempL1 != nil && tempL2 != nil {
		if tempL1.Val != tempL2.Val {
			return false
		}
		tempL1 = tempL1.Next
		tempL2 = tempL2.Next
	}

	// If one of lists didn't finish the elements they are not equal
	if tempL1 != nil || tempL2 != nil {
		return false
	}

	return true
}

// Reverse - returns head to reversed linked list
// Doesn't edit the passed lists.
func Reverse(list *ListNode) *ListNode {
	var prev *ListNode
	cur := list
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

package middleOfLinkedList

// URL 		  - https://leetcode.com/problems/middle-of-the-linked-list/
// Difficulty - Easy

/*
 * Given the head of a singly linked list, return the middle node of the linked list.
 *
 * If there are two middle nodes, return the second middle node.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func middleNode(head *ListNode) *ListNode {
	// If there is just one node
	if head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

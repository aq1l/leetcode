package middleOfLinkedList

import "leetcode/utils/linkedList"

// URL 		  - https://leetcode.com/problems/middle-of-the-linked-list/
// Difficulty - Easy

/*
 * Given the head of a singly linked list, return the middle node of the linked list.
 *
 * If there are two middle nodes, return the second middle node.
 */

func middleNode(head *linkedList.ListNode) *linkedList.ListNode {
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

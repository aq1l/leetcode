package addTwoNumbers

import "leetcode/utils/linkedList"

// URL 		  - https://leetcode.com/problems/add-two-numbers/
// Difficulty - Medium

/*
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order, and each of their nodes contains a single digit.
 * Add the two numbers and return the sum as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 */

func addTwoNumbers(l1 *linkedList.ListNode, l2 *linkedList.ListNode) *linkedList.ListNode {
	answer := &linkedList.ListNode{}
	next := answer
	reminder := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + reminder
		reminder = 0

		if val >= 10 {
			val %= 10
			reminder = 1
		}
		next.Val = val
		if l1.Next != nil || l2.Next != nil {
			next.Next = &linkedList.ListNode{}
			next = next.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		next.Val = l1.Val + reminder
		reminder = 0

		if next.Val >= 10 {
			next.Val %= 10
			reminder = 1
		}
		if l1.Next != nil {
			next.Next = &linkedList.ListNode{}
			next = next.Next
		}
		l1 = l1.Next
	}

	for l2 != nil {
		next.Val = l2.Val + reminder
		reminder = 0

		if next.Val >= 10 {
			next.Val %= 10
			reminder = 1
		}
		if l2.Next != nil {
			next.Next = &linkedList.ListNode{}
			next = next.Next
		}
		l2 = l2.Next
	}

	if reminder > 0 {
		next.Next = &linkedList.ListNode{}
		next = next.Next
		next.Val = reminder
	}

	return answer
}

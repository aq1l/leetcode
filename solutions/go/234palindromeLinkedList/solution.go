package palindromeLinkedList

import "leetcode/utils/go/linkedList"

// URL 		  - https://leetcode.com/problems/palindrome-linked-list/
// Difficulty - Easy

/*
 * Given the head of a singly linked list, return true if it is a palindrome.
 */

// Converting to slice then check if palindrome
func isPalindromeWithSlice(head *linkedList.ListNode) bool {
	temp := head
	var values []int
	for temp != nil {
		values = append(values, temp.Val)
		temp = temp.Next
	}

	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-i-1] {
			return false
		}
	}

	return true
}

// 1. Find midpoint
// 2. Reverse the half
// 3. Compare values
func isPalindrome(head *linkedList.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	revers := linkedList.Reverse(slow)
	for revers != nil {
		if head.Val != revers.Val {
			return false
		}
		head = head.Next
		revers = revers.Next
	}

	return true
}

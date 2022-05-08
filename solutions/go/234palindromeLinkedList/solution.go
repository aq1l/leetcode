package palindromeLinkedList

// URL 		  - https://leetcode.com/problems/palindrome-linked-list/
// Difficulty - Easy

/*
 * Given the head of a singly linked list, return true if it is a palindrome.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromSlice(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head, prev := &ListNode{}, &ListNode{}
	head.Val = data[0]
	prev = head

	data = data[1:]
	for i := range data {
		next := &ListNode{}
		next.Val = data[i]
		prev.Next = next
		prev = next
	}

	return head
}

// Converting to slice then check if palindrome
func isPalindromeWithSlice(head *ListNode) bool {
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
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	revers := reverseList(slow)
	for revers != nil {
		if head.Val != revers.Val {
			return false
		}
		head = head.Next
		revers = revers.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

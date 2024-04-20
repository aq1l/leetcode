package mergeTwoSortedList

import "leetcode/utils/linkedList"

/*
 * You are given the heads of two sorted linked lists list1 and list2.
 *
 * Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
 *
 * Return the head of the merged linked list.
 */

func mergeTwoLists(list1 *linkedList.ListNode, list2 *linkedList.ListNode) *linkedList.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	newList := &linkedList.ListNode{}
	head := newList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newList.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			newList.Val = list2.Val
			list2 = list2.Next
		}

		newList.Next = &linkedList.ListNode{}
		newList = newList.Next
	}

	for list1 != nil {
		newList.Val = list1.Val
		list1 = list1.Next

		if list1 != nil {
			newList.Next = &linkedList.ListNode{}
			newList = newList.Next
		}
	}

	for list2 != nil {
		newList.Val = list2.Val
		list2 = list2.Next

		if list2 != nil {
			newList.Next = &linkedList.ListNode{}
			newList = newList.Next
		}
	}

	return head
}

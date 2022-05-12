package palindromeLinkedList

import (
	"leetcode/utils/go/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeWithSlice(t *testing.T) {
	var testData = []struct {
		input    *linkedList.ListNode
		expected bool
	}{
		{linkedList.NewListFromSlice([]int{1, 2, 2, 1}), true},
		{linkedList.NewListFromSlice([]int{1, 2}), false},
		{linkedList.NewListFromSlice([]int{1, 2, 3, 2, 1}), true},
		{linkedList.NewListFromSlice([]int{1}), true},
	}

	for _, tt := range testData {
		actual := isPalindromeWithSlice(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestIsPalindrome(t *testing.T) {
	var testData = []struct {
		input    *linkedList.ListNode
		expected bool
	}{
		{linkedList.NewListFromSlice([]int{1, 2, 2, 1}), true},
		{linkedList.NewListFromSlice([]int{1, 2}), false},
		{linkedList.NewListFromSlice([]int{1, 2, 3, 2, 1}), true},
		{linkedList.NewListFromSlice([]int{1}), true},
	}

	for _, tt := range testData {
		actual := isPalindrome(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

package palindromeLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeWithSlice(t *testing.T) {
	var testData = []struct {
		input    *ListNode
		expected bool
	}{
		{NewListFromSlice([]int{1, 2, 2, 1}), true},
		{NewListFromSlice([]int{1, 2}), false},
		{NewListFromSlice([]int{1, 2, 3, 2, 1}), true},
		{NewListFromSlice([]int{1}), true},
	}

	for _, tt := range testData {
		actual := isPalindromeWithSlice(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestIsPalindrome(t *testing.T) {
	var testData = []struct {
		input    *ListNode
		expected bool
	}{
		{NewListFromSlice([]int{1, 2, 2, 1}), true},
		{NewListFromSlice([]int{1, 2}), false},
		{NewListFromSlice([]int{1, 2, 3, 2, 1}), true},
		{NewListFromSlice([]int{1}), true},
	}

	for _, tt := range testData {
		actual := isPalindrome(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

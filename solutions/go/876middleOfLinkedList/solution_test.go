package middleOfLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	var testData = []struct {
		head     *ListNode
		expected int
	}{
		{NewListFromSlice([]int{1, 2, 3, 4}), 3},
		{NewListFromSlice([]int{1, 2}), 2},
		{NewListFromSlice([]int{1, 2, 3, 4, 5}), 3},
		{NewListFromSlice([]int{1}), 1},
	}

	for _, tt := range testData {
		actual := middleNode(tt.head)
		assert.Equal(t, tt.expected, actual.Val)
	}
}

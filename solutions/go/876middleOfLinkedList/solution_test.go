package middleOfLinkedList

import (
	"leetcode/utils/go/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	var testData = []struct {
		head     *linkedList.ListNode
		expected int
	}{
		{linkedList.NewListFromSlice([]int{1, 2, 3, 4}), 3},
		{linkedList.NewListFromSlice([]int{1, 2}), 2},
		{linkedList.NewListFromSlice([]int{1, 2, 3, 4, 5}), 3},
		{linkedList.NewListFromSlice([]int{1}), 1},
	}

	for _, tt := range testData {
		actual := middleNode(tt.head)
		assert.Equal(t, tt.expected, actual.Val)
	}
}

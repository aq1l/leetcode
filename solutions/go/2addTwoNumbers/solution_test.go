package addTwoNumbers

import (
	"leetcode/utils/go/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	var testData = []struct {
		l1       *linkedList.ListNode
		l2       *linkedList.ListNode
		expected *linkedList.ListNode
	}{
		{
			linkedList.NewListFromSlice([]int{2, 4, 3}),
			linkedList.NewListFromSlice([]int{5, 6, 4}),
			linkedList.NewListFromSlice([]int{7, 0, 8}),
		},
		{
			linkedList.NewListFromSlice([]int{0}),
			linkedList.NewListFromSlice([]int{0}),
			linkedList.NewListFromSlice([]int{0}),
		},
		{
			linkedList.NewListFromSlice([]int{9, 9, 9, 9, 9, 9, 9}),
			linkedList.NewListFromSlice([]int{9, 9, 9, 9}),
			linkedList.NewListFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tt := range testData {
		actual := addTwoNumbers(tt.l1, tt.l2)
		assert.True(t, linkedList.CompareLists(tt.expected, actual))
	}
}

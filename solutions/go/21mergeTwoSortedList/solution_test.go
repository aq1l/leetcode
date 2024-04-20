package mergeTwoSortedList

import (
	"leetcode/utils/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoList(t *testing.T) {
	var testData = []struct {
		list1    *linkedList.ListNode
		list2    *linkedList.ListNode
		expected *linkedList.ListNode
	}{
		{
			linkedList.NewListFromSlice([]int{1, 2, 4}),
			linkedList.NewListFromSlice([]int{1, 3, 4}),
			linkedList.NewListFromSlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			linkedList.NewListFromSlice([]int{}),
			linkedList.NewListFromSlice([]int{}),
			linkedList.NewListFromSlice([]int{}),
		},
		{
			linkedList.NewListFromSlice([]int{}),
			linkedList.NewListFromSlice([]int{0}),
			linkedList.NewListFromSlice([]int{0}),
		},
	}

	for _, tt := range testData {
		actual := mergeTwoLists(tt.list1, tt.list2)
		assert.True(t, linkedList.CompareLists(tt.expected, actual),
			"Inputs: \n%s\n%s\nExpected: %s\nActual: %s", tt.list1, tt.list2, tt.expected, actual)
	}
}

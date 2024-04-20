package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList(t *testing.T) {
	t.Run("TestCompareLists", func(t *testing.T) {
		t.Run("ListValuesAreEqual", func(t *testing.T) {
			l1 := NewListFromSlice([]int{1, 2, 3, 4, 5})
			l2 := NewListFromSlice([]int{1, 2, 3, 4, 5})

			assert.True(t, CompareLists(l1, l2))
		})

		t.Run("ListValuesAreNotEqual", func(t *testing.T) {
			l1 := NewListFromSlice([]int{1, 2, 3, 4, 4})
			l2 := NewListFromSlice([]int{1, 2, 3, 4, 5})

			assert.False(t, CompareLists(l1, l2))
		})

		t.Run("ListLengthAreNotEqual", func(t *testing.T) {
			l1 := NewListFromSlice([]int{1, 2, 3, 4})
			l2 := NewListFromSlice([]int{1, 2, 3, 4, 5})

			assert.False(t, CompareLists(l1, l2))
		})
	})

	t.Run("TestReverse", func(t *testing.T) {
		l1 := NewListFromSlice([]int{1, 2, 3, 4, 5})
		expected := NewListFromSlice([]int{5, 4, 3, 2, 1})
		l2 := Reverse(l1)
		assert.True(t, CompareLists(expected, l2))
	})
}

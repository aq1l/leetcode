package richestCustomerWealth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumWealth(t *testing.T) {
	var testData = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{1, 2, 3},
			{3, 2, 1},
		}, 6},
		{[][]int{
			{1, 5},
			{7, 3},
			{3, 5},
		}, 10},
		{[][]int{
			{2, 8, 7},
			{7, 1, 3},
			{1, 9, 5},
		}, 17},
	}

	for _, tt := range testData {
		actual := maximumWealth(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

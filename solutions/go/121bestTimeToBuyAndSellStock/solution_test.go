package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testData := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 1, 5, 3, 6, 4, 10}, 9},
	}

	for _, tt := range testData {
		actual := maxProfit(tt.prices)
		assert.Equal(t, tt.expected, actual, "Input: %v\nExpected: %v\nActual: %v", tt.prices, tt.expected, actual)
	}
}

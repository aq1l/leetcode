package kWeakestRowsInMatrix

import (
	"testing"

	//"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/assert"
)

func TestKWeakestRows(t *testing.T) {
	var testData = []struct {
		mat      [][]int
		k        int
		expected []int
	}{
		{[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}, 3, []int{2, 0, 3}},
		{[][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		}, 2, []int{0, 2}},
	}

	for _, tt := range testData {
		actual := kWeakestRows(tt.mat, tt.k)
		assert.ElementsMatch(t, tt.expected, actual)
	}
}

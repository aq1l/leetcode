package steps2ReduceNumber2Zero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSteps(t *testing.T) {
	var testData = []struct {
		input    int
		expected int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
	}

	for _, tt := range testData {
		actual := numberOfSteps(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

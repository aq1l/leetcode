package validParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	var testData = []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range testData {
		actual := isValid(tt.s)
		assert.Equal(t, tt.expected, actual, "Input: %s\nExpected: %v\nActual: %v", tt.s, tt.expected, actual)
	}
}

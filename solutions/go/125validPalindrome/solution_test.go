package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testData := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"a1,2!3:321a", true},
		{"Marge, let's \"[went].\" I await {news} telegram.", true},
	}

	for _, tt := range testData {
		actual := isPalindrome(tt.s)
		assert.Equal(t, tt.expected, actual, "Input: %v\nExpected: %v\nActual: %v", tt.s, tt.expected, actual)
	}
}

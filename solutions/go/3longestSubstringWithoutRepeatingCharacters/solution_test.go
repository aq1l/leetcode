package longestSubstringWithoutRepeatingCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var testData = []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"a", 1},
		{"", 0},
		{"abcd", 4},
		{"dvdf", 3},
	}

	for _, tt := range testData {
		actual := lengthOfLongestSubstring(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

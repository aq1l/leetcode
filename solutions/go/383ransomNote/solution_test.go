package ransomNote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeWithSlice(t *testing.T) {
	var testData = []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aba", true},
	}

	for _, tt := range testData {
		actual := canConstruct(tt.ransomNote, tt.magazine)
		assert.Equal(t, tt.expected, actual)
	}
}

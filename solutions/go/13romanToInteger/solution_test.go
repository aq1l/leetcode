package roman_to_int

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	var testData = []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range testData {
		actual := romanToInt(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}

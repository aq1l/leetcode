package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuz(t *testing.T) {
	var testData = []struct {
		input    int
		expected []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, tt := range testData {
		actual := fizzBuzz(tt.input)
		assert.ElementsMatch(t, tt.expected, actual)
	}
}

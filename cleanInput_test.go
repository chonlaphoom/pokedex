package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "h w",
			expected: []string{"h", "w"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("Input: %s, Expected: %v, Actual: %v\n", c.input, c.expected, actual)
		for i := range cases {

			if actual[i] != c.expected[i] {
				t.Errorf("Expected %v, but got %v", c.expected, actual)
			}
		}
	}
}

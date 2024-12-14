package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "testing    whether,   this works      ",
			expected: []string{"testing", "whether,", "this", "works"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Fatalf("actual and expected have different lengths")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %s and word %s at index %d do not match", word, expectedWord, i)
				t.Fail()
			}
		}

	}
}

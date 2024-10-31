package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello", 
				"world",
			},
		},
		{
			input: "THIS is a text",
			expected: []string{
				"this",
				"is",
				"a",
				"text",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal")
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v not equal to %v", actualWord, expectedWord)
			}
		}
	}
}
package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"snow dog sun", "sun dog snow"},
		{"hello world", "world hello"},
		{"one", "one"},
		{"", ""},
		{"a b c d e", "e d c b a"},
		{"  leading spaces", "spaces leading"},
		{"trailing spaces  ", "spaces trailing"},
		{"  both  sides  ", "sides both"},
	}

	for _, test := range tests {
		result := ReverseWords(test.input)
		if result != test.expected {
			t.Errorf("ReverseWords(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

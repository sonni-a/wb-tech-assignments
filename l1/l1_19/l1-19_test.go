package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"главрыба", "абырвалг"},
		{"hello", "olleh"},
		{"12345", "54321"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

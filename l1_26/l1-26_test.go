package main

import "testing"

func TestIsAllUnique(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
		{"", true},
		{"a", true},
		{"Aa", false},
	}

	for _, tt := range tests {
		result := isAllUnique(tt.input)
		if result != tt.expected {
			t.Errorf("for %q got %v; want %v", tt.input, result, tt.expected)
		}
	}
}

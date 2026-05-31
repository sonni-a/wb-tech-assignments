package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "found in middle",
			nums:     []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "found at beginning",
			nums:     []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "found at end",
			nums:     []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "not found",
			nums:     []int{1, 3, 5, 7, 9},
			target:   6,
			expected: -1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "one element found",
			nums:     []int{10},
			target:   10,
			expected: 0,
		},
		{
			name:     "one element not found",
			nums:     []int{10},
			target:   5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := binarySearch(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

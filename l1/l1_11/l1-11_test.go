package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "basic case",
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "no intersection",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: []int{},
		},
		{
			name:     "with duplicates",
			nums1:    []int{1, 2, 2, 3},
			nums2:    []int{2, 2, 4},
			expected: []int{2},
		},
		{
			name:     "empty input",
			nums1:    []int{},
			nums2:    []int{1, 2},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersection(tt.nums1, tt.nums2)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

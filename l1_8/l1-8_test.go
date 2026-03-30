package main

import "testing"

func TestSetBit(t *testing.T) {
	tests := []struct {
		name     string
		num      int64
		i        int
		val      int
		expected int64
		hasError bool
	}{
		{
			name:     "set bit to 0",
			num:      5,
			i:        1,
			val:      0,
			expected: 5 &^ (1 << 1),
		},
		{
			name:     "set bit to 1",
			num:      5,
			i:        1,
			val:      1,
			expected: 7,
		},
		{
			name:     "clear bit",
			num:      5,
			i:        0,
			val:      0,
			expected: 4,
		},
		{
			name:     "invalid value",
			num:      5,
			i:        1,
			val:      2,
			expected: 5,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SetBit(tt.num, tt.i, tt.val)

			if tt.hasError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

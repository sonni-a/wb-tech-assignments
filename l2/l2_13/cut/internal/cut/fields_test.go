package cut_test

import (
	"testing"

	"cut/internal/cut"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		spec  string
		want  []int
		isErr bool
	}{
		{"1", []int{1}, false},
		{"1,3", []int{1, 3}, false},
		{"3-5", []int{3, 4, 5}, false},
		{"1,3-5,7", []int{1, 3, 4, 5, 7}, false},
		{"", nil, true},
		{"1,", nil, true},
		{"abc", nil, true},
		{"1-", nil, true},
		{"-5", nil, true},
		{"5-3", nil, true},
		{"0", nil, true},
	}

	for _, tt := range tests {
		got, err := cut.ParseFields(tt.spec)
		if tt.isErr {
			if err == nil {
				t.Fatalf("ParseFields(%q): expected error", tt.spec)
			}
			continue
		}
		if err != nil {
			t.Fatalf("ParseFields(%q): %v", tt.spec, err)
		}
		if len(got) != len(tt.want) {
			t.Fatalf("ParseFields(%q) = %v, want %v", tt.spec, got, tt.want)
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Fatalf("ParseFields(%q) = %v, want %v", tt.spec, got, tt.want)
			}
		}
	}
}

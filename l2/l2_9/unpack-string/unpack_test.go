package main

import "testing"

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "basic case",
			input: "a4bc2d5e",
			want:  "aaaabccddddde",
		},
		{
			name:  "no digits",
			input: "abcd",
			want:  "abcd",
		},
		{
			name:    "only digits invalid",
			input:   "45",
			want:    "",
			wantErr: true,
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "escaped digits",
			input: `qwe\4\5`,
			want:  "qwe45",
		},
		{
			name:  "escaped + repeat",
			input: `qwe\45`,
			want:  "qwe44444",
		},
		{
			name:    "invalid escape",
			input:   "abc\\",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackString(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

package urlpath

import (
	"path/filepath"
	"testing"
)

func TestLocalPath(t *testing.T) {
	tests := []struct {
		raw  string
		want string
	}{
		{"https://example.com/", filepath.Join("example.com", "index.html")},
		{"https://example.com/about", filepath.Join("example.com", "about", "index.html")},
		{"https://example.com/about/", filepath.Join("example.com", "about", "index.html")},
		{"https://example.com/css/style.css", filepath.Join("example.com", "css", "style.css")},
	}

	for _, tt := range tests {
		got, err := LocalPath(tt.raw)
		if err != nil {
			t.Fatalf("LocalPath(%q): %v", tt.raw, err)
		}
		if got != tt.want {
			t.Errorf("LocalPath(%q) = %q, want %q", tt.raw, got, tt.want)
		}
	}
}

func TestRelLink(t *testing.T) {
	from := filepath.Join("example.com", "docs", "page.html")
	to := filepath.Join("example.com", "css", "style.css")
	got := RelLink(from, to)
	want := "../css/style.css"
	if got != want {
		t.Errorf("RelLink() = %q, want %q", got, want)
	}
}

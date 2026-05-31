package urlutil

import "testing"

func TestResolve(t *testing.T) {
	base := "https://example.com/docs/page.html"

	tests := []struct {
		ref  string
		want string
	}{
		{"/style.css", "https://example.com/style.css"},
		{"../img/logo.png", "https://example.com/img/logo.png"},
		{"https://other.com/x", "https://other.com/x"},
		{"#section", ""},
		{"mailto:a@b.c", ""},
		{"javascript:void(0)", ""},
		{"data:image/png;base64,abc", ""},
	}

	for _, tt := range tests {
		got := Resolve(base, tt.ref)
		if got != tt.want {
			t.Errorf("Resolve(%q, %q) = %q, want %q", base, tt.ref, got, tt.want)
		}
	}
}

func TestSameHost(t *testing.T) {
	if !SameHost("https://Example.com/a", "http://example.com/b") {
		t.Fatal("expected same host")
	}
	if SameHost("https://example.com/a", "https://other.com/b") {
		t.Fatal("expected different hosts")
	}
}

func TestLooksLikeHTMLPath(t *testing.T) {
	if !LooksLikeHTMLPath("https://example.com/") {
		t.Fatal("root should look like HTML")
	}
	if !LooksLikeHTMLPath("https://example.com/about.html") {
		t.Fatal(".html should look like HTML")
	}
	if LooksLikeHTMLPath("https://example.com/app.js") {
		t.Fatal(".js should not look like HTML")
	}
}

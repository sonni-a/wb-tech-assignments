package crawler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"wget/internal/downloader"
)

func TestMirrorExampleCom(t *testing.T) {
	out := t.TempDir()
	client := downloader.New(30 * time.Second)
	m, err := New(Config{
		StartURL: "https://example.com",
		Output:   out,
		Depth:    0,
		Workers:  2,
	}, client)
	if err != nil {
		t.Fatal(err)
	}

	if err := m.Run(); err != nil {
		t.Fatal(err)
	}

	indexPath := filepath.Join(out, "example.com", "index.html")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("index.html missing: %v", err)
	}

	body := string(data)
	if !strings.Contains(body, "Example Domain") {
		t.Fatalf("unexpected page content: %s", body)
	}
	if strings.Contains(body, "https://example.com") {
		t.Fatalf("remote links not rewritten: %s", body)
	}
}

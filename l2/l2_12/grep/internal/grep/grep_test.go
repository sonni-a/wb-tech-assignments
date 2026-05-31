package grep_test

import (
	"bytes"
	"strings"
	"testing"

	"grep/internal/grep"
)

func TestRunBasicMatch(t *testing.T) {
	lines := []string{"hello", "world", "help"}
	m, err := grep.NewMatcher("hel", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	count, found := grep.Run(lines, m, grep.Options{}, &out)
	if !found || count != 2 {
		t.Fatalf("expected 2 matches, got count=%d found=%v", count, found)
	}

	got := strings.TrimSpace(out.String())
	want := "hello\nhelp"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunIgnoreCase(t *testing.T) {
	lines := []string{"Hello", "WORLD", "help"}
	m, err := grep.NewMatcher("hello", grep.Options{IgnoreCase: true})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{IgnoreCase: true}, &out)

	if strings.TrimSpace(out.String()) != "Hello" {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunInvert(t *testing.T) {
	lines := []string{"aa", "bb", "cc"}
	m, err := grep.NewMatcher("bb", grep.Options{Invert: true})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{Invert: true}, &out)

	got := strings.TrimSpace(out.String())
	want := "aa\ncc"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunFixedString(t *testing.T) {
	lines := []string{"a.b", "abc", "x"}
	m, err := grep.NewMatcher("a.b", grep.Options{FixedString: true})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{FixedString: true}, &out)

	if strings.TrimSpace(out.String()) != "a.b" {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunCount(t *testing.T) {
	lines := []string{"one", "two", "one", "three"}
	m, err := grep.NewMatcher("one", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{Count: true}, &out)

	if strings.TrimSpace(out.String()) != "2" {
		t.Fatalf("unexpected count: %q", out.String())
	}
}

func TestRunLineNumber(t *testing.T) {
	lines := []string{"first", "match", "last"}
	m, err := grep.NewMatcher("match", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{LineNumber: true}, &out)

	if strings.TrimSpace(out.String()) != "2:match" {
		t.Fatalf("unexpected output: %q", out.String())
	}
}

func TestRunContext(t *testing.T) {
	lines := []string{"1", "2", "hit", "4", "5", "hit", "7"}
	m, err := grep.NewMatcher("hit", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	opts := grep.Options{BeforeContext: 1, AfterContext: 1, LineNumber: true}
	var out bytes.Buffer
	grep.Run(lines, m, opts, &out)

	got := strings.TrimSpace(out.String())
	want := "2-2\n3:hit\n4-4\n5-5\n6:hit\n7-7"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunContextSeparator(t *testing.T) {
	lines := []string{"1", "hit", "3", "4", "5", "hit", "7"}
	m, err := grep.NewMatcher("hit", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	opts := grep.Options{BeforeContext: 0, AfterContext: 0, LineNumber: true}
	var out bytes.Buffer
	grep.Run(lines, m, opts, &out)

	got := strings.TrimSpace(out.String())
	want := "2:hit\n6:hit"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}

	opts = grep.Options{BeforeContext: 1, AfterContext: 1, LineNumber: true}
	out.Reset()
	grep.Run(lines, m, opts, &out)

	got = strings.TrimSpace(out.String())
	want = "1-1\n2:hit\n3-3\n--\n5-5\n6:hit\n7-7"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunContextMerge(t *testing.T) {
	lines := []string{"a", "hit", "c", "hit", "e"}
	m, err := grep.NewMatcher("hit", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	opts := grep.Options{BeforeContext: 1, AfterContext: 1}
	var out bytes.Buffer
	grep.Run(lines, m, opts, &out)

	got := strings.TrimSpace(out.String())
	want := "a\nhit\nc\nhit\ne"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunRegex(t *testing.T) {
	lines := []string{"foo123", "bar", "foo456"}
	m, err := grep.NewMatcher(`foo\d+`, grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	grep.Run(lines, m, grep.Options{}, &out)

	got := strings.TrimSpace(out.String())
	want := "foo123\nfoo456"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunNoMatch(t *testing.T) {
	lines := []string{"a", "b", "c"}
	m, err := grep.NewMatcher("z", grep.Options{})
	if err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	_, found := grep.Run(lines, m, grep.Options{}, &out)
	if found {
		t.Fatal("expected no matches")
	}
	if out.Len() != 0 {
		t.Fatalf("expected empty output, got %q", out.String())
	}
}

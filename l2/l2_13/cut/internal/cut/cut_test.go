package cut_test

import (
	"bytes"
	"strings"
	"testing"

	"cut/internal/cut"
)

func TestProcessLineBasic(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{1, 3},
		Delimiter: "\t",
	}

	out, ok := cut.ProcessLine("a\tb\tc", opts)
	if !ok || out != "a\tc" {
		t.Fatalf("got %q, ok=%v; want %q", out, ok, "a\tc")
	}
}

func TestProcessLineCustomDelimiter(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{2},
		Delimiter: ",",
	}

	out, ok := cut.ProcessLine("one,two,three", opts)
	if !ok || out != "two" {
		t.Fatalf("got %q, ok=%v; want %q", out, ok, "two")
	}
}

func TestProcessLineOutOfBounds(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{1, 5},
		Delimiter: "\t",
	}

	out, ok := cut.ProcessLine("a\tb", opts)
	if !ok || out != "a" {
		t.Fatalf("got %q, ok=%v; want %q", out, ok, "a")
	}
}

func TestProcessLineNoDelimiterWithoutSeparated(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{1},
		Delimiter: "\t",
	}

	out, ok := cut.ProcessLine("hello", opts)
	if !ok || out != "hello" {
		t.Fatalf("got %q, ok=%v; want %q", out, ok, "hello")
	}
}

func TestProcessLineSeparated(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{1},
		Delimiter: "\t",
		Separated: true,
	}

	out, ok := cut.ProcessLine("hello", opts)
	if ok {
		t.Fatalf("expected line to be skipped, got %q", out)
	}
}

func TestProcessLineFieldOrder(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{3, 1},
		Delimiter: "\t",
	}

	out, ok := cut.ProcessLine("a\tb\tc", opts)
	if !ok || out != "c\ta" {
		t.Fatalf("got %q, ok=%v; want %q", out, ok, "c\ta")
	}
}

func TestProcessLineEmptyField(t *testing.T) {
	opts := cut.Options{
		Fields:    []int{2},
		Delimiter: "\t",
	}

	out, ok := cut.ProcessLine("a\t\tc", opts)
	if !ok || out != "" {
		t.Fatalf("got %q, ok=%v; want empty string", out, ok)
	}
}

func TestRun(t *testing.T) {
	input := "a\tb\tc\nd\te\n"
	opts := cut.Options{
		Fields:    []int{1, 3},
		Delimiter: "\t",
	}

	var out bytes.Buffer
	if err := cut.Run(strings.NewReader(input), &out, opts); err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(out.String())
	want := "a\tc\nd"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunSeparated(t *testing.T) {
	input := "a\tb\nplain\n"
	opts := cut.Options{
		Fields:    []int{1},
		Delimiter: "\t",
		Separated: true,
	}

	var out bytes.Buffer
	if err := cut.Run(strings.NewReader(input), &out, opts); err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(out.String())
	want := "a"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

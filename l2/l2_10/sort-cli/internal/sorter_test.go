package sorter

import "testing"

func TestNumericSort(t *testing.T) {
	lines := []string{"10", "2", "1"}
	opts := Options{Numeric: true}
	Sort(lines, opts)

	want := []string{"1", "2", "10"}
	assertLines(t, lines, want)
}

func TestNumericReverse(t *testing.T) {
	lines := []string{"1", "3", "2"}
	opts := Options{Numeric: true, Reverse: true}
	Sort(lines, opts)

	want := []string{"3", "2", "1"}
	assertLines(t, lines, want)
}

func TestColumnSort(t *testing.T) {
	lines := []string{"b\ta", "c\tb", "a\tc"}
	opts := Options{Column: 2}
	Sort(lines, opts)

	want := []string{"b\ta", "c\tb", "a\tc"}
	assertLines(t, lines, want)
}

func TestUniqueByColumn(t *testing.T) {
	lines := []string{"1\tx", "2\tx", "3\ty"}
	opts := Options{Column: 2}
	Sort(lines, opts)
	lines = Unique(lines, opts)

	want := []string{"1\tx", "3\ty"}
	assertLines(t, lines, want)
}

func TestMonthSort(t *testing.T) {
	lines := []string{"Mar", "Jan", "Feb"}
	opts := Options{Month: true}
	Sort(lines, opts)

	want := []string{"Jan", "Feb", "Mar"}
	assertLines(t, lines, want)
}

func TestHumanSort(t *testing.T) {
	lines := []string{"10M", "2K", "500"}
	opts := Options{Human: true}
	Sort(lines, opts)

	want := []string{"500", "2K", "10M"}
	assertLines(t, lines, want)
}

func TestIsSorted(t *testing.T) {
	lines := []string{"a", "b", "c"}
	if !IsSorted(lines, Options{}) {
		t.Fatal("expected sorted lines")
	}

	lines = []string{"b", "a"}
	if IsSorted(lines, Options{}) {
		t.Fatal("expected unsorted lines")
	}
}

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

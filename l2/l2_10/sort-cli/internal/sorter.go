// Package sorter provides sorting logic for the sort CLI utility.
package sorter

import "sort"

// Options holds command-line flags that control sorting behavior.
type Options struct {
	Column   int
	Numeric  bool
	Reverse  bool
	Unique   bool
	Month    bool
	IgnoreTB bool
	Check    bool
	Human    bool
}

// Sort sorts lines in place using the given options.
func Sort(lines []string, opts Options) {
	sort.Slice(lines, func(i, j int) bool {
		return less(lines[i], lines[j], opts)
	})
}

// Unique removes consecutive duplicate lines by sort key after sorting.
func Unique(lines []string, opts Options) []string {
	if len(lines) == 0 {
		return lines
	}

	result := []string{lines[0]}
	prevKey := extractKey(lines[0], opts)

	for i := 1; i < len(lines); i++ {
		key := extractKey(lines[i], opts)
		if key != prevKey {
			result = append(result, lines[i])
			prevKey = key
		}
	}

	return result
}

// IsSorted reports whether lines are already sorted according to opts.
func IsSorted(lines []string, opts Options) bool {
	for i := 1; i < len(lines); i++ {
		if less(lines[i], lines[i-1], opts) {
			return false
		}
	}
	return true
}

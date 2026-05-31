package grep

import (
	"fmt"
	"io"
	"strings"
)

type lineRange struct {
	start int
	end   int
}

// Run filters lines from input and writes matching output to out.
// It returns the number of matching lines and whether any line matched.
func Run(lines []string, matcher *Matcher, opts Options, out io.Writer) (int, bool) {
	matched := make([]bool, len(lines))
	count := 0

	for i, line := range lines {
		if matcher.Match(line) {
			matched[i] = true
			count++
		}
	}

	if opts.Count {
		fmt.Fprintln(out, count)
		return count, count > 0
	}

	if opts.BeforeContext > 0 || opts.AfterContext > 0 {
		printWithContext(lines, matched, opts, out)
		return count, count > 0
	}

	for i, line := range lines {
		if matched[i] {
			writeLine(out, line, i+1, opts.LineNumber, true)
		}
	}

	return count, count > 0
}

func printWithContext(lines []string, matched []bool, opts Options, out io.Writer) {
	ranges := buildContextRanges(len(lines), matched, opts.BeforeContext, opts.AfterContext)
	for g, r := range ranges {
		if g > 0 {
			fmt.Fprintln(out, "--")
		}
		for i := r.start; i <= r.end; i++ {
			writeLine(out, lines[i], i+1, opts.LineNumber, matched[i])
		}
	}
}

func buildContextRanges(n int, matched []bool, before, after int) []lineRange {
	var ranges []lineRange

	for i := 0; i < n; i++ {
		if !matched[i] {
			continue
		}
		start := i - before
		if start < 0 {
			start = 0
		}
		end := i + after
		if end >= n {
			end = n - 1
		}
		ranges = append(ranges, lineRange{start: start, end: end})
	}

	return mergeRanges(ranges)
}

func mergeRanges(ranges []lineRange) []lineRange {
	if len(ranges) == 0 {
		return nil
	}

	merged := []lineRange{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.start <= last.end+1 {
			if r.end > last.end {
				last.end = r.end
			}
			continue
		}
		merged = append(merged, r)
	}

	return merged
}

func writeLine(out io.Writer, line string, num int, showNum, isMatch bool) {
	if !showNum {
		fmt.Fprintln(out, line)
		return
	}

	sep := "-"
	if isMatch {
		sep = ":"
	}
	fmt.Fprintf(out, "%d%s%s\n", num, sep, line)
}

// FormatLine returns a single output line with optional line number prefix.
func FormatLine(line string, num int, showNum, isMatch bool) string {
	var b strings.Builder
	if showNum {
		sep := "-"
		if isMatch {
			sep = ":"
		}
		fmt.Fprintf(&b, "%d%s", num, sep)
	}
	b.WriteString(line)
	return b.String()
}

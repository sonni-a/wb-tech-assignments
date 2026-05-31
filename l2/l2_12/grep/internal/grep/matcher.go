package grep

import (
	"regexp"
	"strings"
)

// Matcher checks whether a line matches the search pattern.
type Matcher struct {
	fixed  string
	re     *regexp.Regexp
	invert bool
	ignore bool
}

// NewMatcher builds a matcher from pattern and options.
func NewMatcher(pattern string, opts Options) (*Matcher, error) {
	m := &Matcher{invert: opts.Invert, ignore: opts.IgnoreCase}

	if opts.FixedString {
		m.fixed = pattern
		return m, nil
	}

	expr := pattern
	if opts.IgnoreCase {
		expr = "(?i)" + pattern
	}

	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	m.re = re
	return m, nil
}

// Match reports whether the line satisfies the filter.
func (m *Matcher) Match(line string) bool {
	matched := m.matchLine(line)
	if m.invert {
		return !matched
	}
	return matched
}

func (m *Matcher) matchLine(line string) bool {
	if m.fixed != "" {
		if m.ignore {
			return strings.Contains(strings.ToLower(line), strings.ToLower(m.fixed))
		}
		return strings.Contains(line, m.fixed)
	}
	return m.re.MatchString(line)
}

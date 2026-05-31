package sorter

import (
	"strconv"
	"strings"
)

var monthOrder = map[string]int{
	"jan": 1, "feb": 2, "mar": 3, "apr": 4,
	"may": 5, "jun": 6, "jul": 7, "aug": 8,
	"sep": 9, "oct": 10, "nov": 11, "dec": 12,
}

func less(a, b string, opts Options) bool {
	ka := extractKey(a, opts)
	kb := extractKey(b, opts)

	var result bool

	switch {
	case opts.Numeric:
		result = toFloat(ka) < toFloat(kb)
	case opts.Human:
		result = parseHuman(ka) < parseHuman(kb)
	case opts.Month:
		result = parseMonth(ka) < parseMonth(kb)
	default:
		result = ka < kb
	}

	if opts.Reverse {
		return !result
	}
	return result
}

func toFloat(s string) float64 {
	v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return v
}

func parseMonth(s string) int {
	s = strings.ToLower(strings.TrimSpace(s))
	return monthOrder[s]
}

func parseHuman(s string) float64 {
	s = strings.ToUpper(strings.TrimSpace(s))

	mult := 1.0

	switch {
	case strings.HasSuffix(s, "K"):
		mult = 1e3
		s = strings.TrimSuffix(s, "K")
	case strings.HasSuffix(s, "M"):
		mult = 1e6
		s = strings.TrimSuffix(s, "M")
	case strings.HasSuffix(s, "G"):
		mult = 1e9
		s = strings.TrimSuffix(s, "G")
	}

	v, _ := strconv.ParseFloat(s, 64)
	return v * mult
}

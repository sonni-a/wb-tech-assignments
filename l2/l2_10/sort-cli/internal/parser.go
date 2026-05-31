package sorter

import "strings"

func extractKey(line string, opts Options) string {
	if opts.IgnoreTB {
		line = strings.TrimRight(line, " \t")
	}

	if opts.Column <= 0 {
		return line
	}

	parts := strings.Split(line, "\t")
	if opts.Column-1 >= len(parts) {
		return ""
	}

	key := parts[opts.Column-1]
	if opts.IgnoreTB {
		key = strings.TrimRight(key, " \t")
	}

	return key
}

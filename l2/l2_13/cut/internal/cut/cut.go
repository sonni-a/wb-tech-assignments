package cut

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Fields    []int
	Delimiter string
	Separated bool
}

func ProcessLine(line string, opts Options) (string, bool) {
	if opts.Separated && !strings.Contains(line, opts.Delimiter) {
		return "", false
	}

	parts := strings.Split(line, opts.Delimiter)

	var out strings.Builder
	first := true
	for _, fieldNum := range opts.Fields {
		if fieldNum < 1 || fieldNum > len(parts) {
			continue
		}
		if !first {
			out.WriteString(opts.Delimiter)
		}
		out.WriteString(parts[fieldNum-1])
		first = false
	}

	return out.String(), true
}

func Run(r io.Reader, w io.Writer, opts Options) error {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		out, ok := ProcessLine(scanner.Text(), opts)
		if !ok {
			continue
		}
		if _, err := fmt.Fprintln(w, out); err != nil {
			return err
		}
	}

	return scanner.Err()
}

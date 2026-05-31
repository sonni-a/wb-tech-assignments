package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"grep/internal/grep"
)

func main() {
	opts, pattern, files, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	matcher, err := grep.NewMatcher(pattern, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	lines, err := readLines(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	_, found := grep.Run(lines, matcher, opts, os.Stdout)
	if !found {
		os.Exit(1)
	}
}

func readLines(files []string) ([]string, error) {
	if len(files) == 0 {
		return readFrom(os.Stdin)
	}

	var lines []string
	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			return nil, err
		}

		fileLines, err := readFrom(f)
		f.Close()
		if err != nil {
			return nil, err
		}

		lines = append(lines, fileLines...)
	}

	return lines, nil
}

func readFrom(r io.Reader) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

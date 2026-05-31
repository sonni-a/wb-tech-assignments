package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	sorter "sort-cli/internal"
)

func main() {
	opts, files, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	lines, err := readLines(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if opts.Check {
		if !sorter.IsSorted(lines, opts) {
			fmt.Println("data is not sorted")
			os.Exit(1)
		}
		return
	}

	sorter.Sort(lines, opts)

	if opts.Unique {
		lines = sorter.Unique(lines, opts)
	}

	for _, line := range lines {
		fmt.Println(line)
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

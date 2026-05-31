package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"grep/internal/grep"
)

func parseArgs() (grep.Options, string, []string, error) {
	opts := grep.Options{}
	var pattern string
	var files []string

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			if pattern == "" && i+1 < len(args) {
				pattern = args[i+1]
				files = append(files, args[i+2:]...)
			} else {
				files = append(files, args[i+1:]...)
			}
			break
		}

		if !strings.HasPrefix(arg, "-") || arg == "-" {
			if pattern == "" {
				pattern = arg
			} else {
				files = append(files, arg)
			}
			continue
		}

		j := 1
		for j < len(arg) {
			switch arg[j] {
			case 'A':
				n, next, err := parseNumber(arg, j, args, &i)
				if err != nil {
					return opts, "", nil, err
				}
				opts.AfterContext = n
				j = next
			case 'B':
				n, next, err := parseNumber(arg, j, args, &i)
				if err != nil {
					return opts, "", nil, err
				}
				opts.BeforeContext = n
				j = next
			case 'C':
				n, next, err := parseNumber(arg, j, args, &i)
				if err != nil {
					return opts, "", nil, err
				}
				opts.BeforeContext = n
				opts.AfterContext = n
				j = next
			case 'c':
				opts.Count = true
				j++
			case 'i':
				opts.IgnoreCase = true
				j++
			case 'v':
				opts.Invert = true
				j++
			case 'F':
				opts.FixedString = true
				j++
			case 'n':
				opts.LineNumber = true
				j++
			default:
				return opts, "", nil, fmt.Errorf("unknown flag: -%c in %s", arg[j], arg)
			}
		}
	}

	if pattern == "" {
		return opts, "", nil, fmt.Errorf("pattern is required")
	}

	return opts, pattern, files, nil
}

func parseNumber(arg string, flagPos int, args []string, idx *int) (int, int, error) {
	rest := arg[flagPos+1:]
	if rest == "" {
		*idx++
		if *idx >= len(args) {
			return 0, len(arg), fmt.Errorf("flag -%c requires a value", arg[flagPos])
		}
		n, err := strconv.Atoi(args[*idx])
		return n, len(arg), err
	}

	end := 0
	for end < len(rest) && rest[end] >= '0' && rest[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, len(arg), fmt.Errorf("invalid value for -%c: %q", arg[flagPos], rest)
	}

	n, err := strconv.Atoi(rest[:end])
	return n, flagPos + 1 + end, err
}

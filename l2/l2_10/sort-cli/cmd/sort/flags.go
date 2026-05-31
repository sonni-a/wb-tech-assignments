package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	sorter "sort-cli/internal"
)

func parseArgs() (sorter.Options, []string, error) {
	opts := sorter.Options{}
	var files []string

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			files = append(files, args[i+1:]...)
			break
		}

		if !strings.HasPrefix(arg, "-") || arg == "-" {
			files = append(files, arg)
			continue
		}

		j := 1
		for j < len(arg) {
			switch arg[j] {
			case 'k':
				col, next, err := parseColumn(arg, j, args, &i)
				if err != nil {
					return opts, nil, err
				}
				opts.Column = col
				j = next
			case 'n':
				opts.Numeric = true
				j++
			case 'r':
				opts.Reverse = true
				j++
			case 'u':
				opts.Unique = true
				j++
			case 'M':
				opts.Month = true
				j++
			case 'b':
				opts.IgnoreTB = true
				j++
			case 'c':
				opts.Check = true
				j++
			case 'h':
				opts.Human = true
				j++
			default:
				return opts, nil, fmt.Errorf("unknown flag: -%c in %s", arg[j], arg)
			}
		}
	}

	return opts, files, nil
}

func parseColumn(arg string, kPos int, args []string, idx *int) (int, int, error) {
	rest := arg[kPos+1:]
	if rest == "" {
		*idx++
		if *idx >= len(args) {
			return 0, len(arg), fmt.Errorf("flag -k requires a value")
		}
		col, err := strconv.Atoi(args[*idx])
		return col, len(arg), err
	}

	if rest[0] == '=' {
		rest = rest[1:]
	}

	end := 0
	for end < len(rest) && rest[end] >= '0' && rest[end] <= '9' {
		end++
	}
	if end == 0 {
		return 0, len(arg), fmt.Errorf("invalid value for -k: %q", rest)
	}

	col, err := strconv.Atoi(rest[:end])
	return col, kPos + 1 + end, err
}

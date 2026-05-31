package main

import (
	"fmt"
	"os"
	"strings"

	"cut/internal/cut"
)

func parseArgs() (cut.Options, error) {
	opts := cut.Options{Delimiter: "\t"}

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if !strings.HasPrefix(arg, "-") || arg == "-" {
			return opts, fmt.Errorf("unexpected argument: %q", arg)
		}

		j := 1
		for j < len(arg) {
			switch arg[j] {
			case 'f':
				spec, next, err := parseFlagValue(arg, j, args, &i)
				if err != nil {
					return opts, err
				}
				fields, err := cut.ParseFields(spec)
				if err != nil {
					return opts, err
				}
				opts.Fields = fields
				j = next
			case 'd':
				val, next, err := parseFlagValue(arg, j, args, &i)
				if err != nil {
					return opts, err
				}
				if val == "" {
					return opts, fmt.Errorf("delimiter must not be empty")
				}
				opts.Delimiter = val
				j = next
			case 's':
				opts.Separated = true
				j++
			default:
				return opts, fmt.Errorf("unknown flag: -%c in %s", arg[j], arg)
			}
		}
	}

	if len(opts.Fields) == 0 {
		return opts, fmt.Errorf("flag -f is required")
	}

	return opts, nil
}

func parseFlagValue(arg string, flagPos int, args []string, idx *int) (string, int, error) {
	rest := arg[flagPos+1:]
	if rest == "" {
		*idx++
		if *idx >= len(args) {
			return "", len(arg), fmt.Errorf("flag -%c requires a value", arg[flagPos])
		}
		return args[*idx], len(arg), nil
	}

	return rest, len(arg), nil
}

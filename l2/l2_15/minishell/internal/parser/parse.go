package parser

import (
	"fmt"
	"os"
	"strings"
)

func expandLine(line string) string {
	return os.ExpandEnv(line)
}

func splitByLogicalOps(line string) ([]string, []string, error) {
	var segments []string
	var ops []string
	var cur strings.Builder

	for i := 0; i < len(line); i++ {
		switch {
		case strings.HasPrefix(line[i:], "||"):
			segments = append(segments, strings.TrimSpace(cur.String()))
			ops = append(ops, "||")
			cur.Reset()
			i++
		case strings.HasPrefix(line[i:], "&&"):
			segments = append(segments, strings.TrimSpace(cur.String()))
			ops = append(ops, "&&")
			cur.Reset()
			i++
		default:
			cur.WriteByte(line[i])
		}
	}
	segments = append(segments, strings.TrimSpace(cur.String()))

	if len(segments) == 0 {
		return nil, nil, nil
	}
	if len(ops) != len(segments)-1 {
		return nil, nil, fmt.Errorf("mismatched logical operators")
	}
	return segments, ops, nil
}

func parseCommandFields(fields []string) (Command, error) {
	var cmd Command

	for i := 0; i < len(fields); i++ {
		switch fields[i] {
		case ">":
			if i+1 >= len(fields) {
				return Command{}, fmt.Errorf("redirect missing filename")
			}
			cmd.Stdout = &Redirect{Path: fields[i+1]}
			cmd.AppendOut = false
			i++
		case ">>":
			if i+1 >= len(fields) {
				return Command{}, fmt.Errorf("redirect missing filename")
			}
			cmd.Stdout = &Redirect{Path: fields[i+1]}
			cmd.AppendOut = true
			i++
		case "<":
			if i+1 >= len(fields) {
				return Command{}, fmt.Errorf("redirect missing filename")
			}
			cmd.Stdin = &Redirect{Path: fields[i+1]}
			i++
		default:
			cmd.Args = append(cmd.Args, fields[i])
		}
	}
	return cmd, nil
}

func parseCommandSegment(segment string) (Command, error) {
	segment = strings.TrimSpace(segment)
	if segment == "" {
		return Command{}, nil
	}
	return parseCommandFields(strings.Fields(segment))
}

func parsePipelineSegment(segment string) (Pipeline, error) {
	var pl Pipeline
	for _, part := range strings.Split(segment, "|") {
		cmd, err := parseCommandSegment(part)
		if err != nil {
			return Pipeline{}, err
		}
		if len(cmd.Args) > 0 || cmd.Stdin != nil || cmd.Stdout != nil {
			pl.Commands = append(pl.Commands, cmd)
		}
	}
	return pl, nil
}

// ParseLine parses a shell input line.
func ParseLine(line string) (*Job, error) {
	line = strings.TrimSpace(expandLine(line))
	if line == "" {
		return nil, nil
	}

	segments, ops, err := splitByLogicalOps(line)
	if err != nil {
		return nil, err
	}
	if len(segments) == 0 {
		return nil, nil
	}

	var j Job
	for _, seg := range segments {
		pl, err := parsePipelineSegment(seg)
		if err != nil {
			return nil, err
		}
		j.Pipelines = append(j.Pipelines, pl)
	}
	j.Ops = ops
	return &j, nil
}

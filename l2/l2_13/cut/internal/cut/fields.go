package cut

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseFields(spec string) ([]int, error) {
	spec = strings.TrimSpace(spec)
	if spec == "" {
		return nil, fmt.Errorf("empty fields specification")
	}

	var fields []int
	for _, part := range strings.Split(spec, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("invalid fields specification: %q", spec)
		}

		if strings.Contains(part, "-") {
			start, end, err := parseRange(part)
			if err != nil {
				return nil, err
			}
			for i := start; i <= end; i++ {
				fields = append(fields, i)
			}
			continue
		}

		n, err := strconv.Atoi(part)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("invalid field number: %q", part)
		}
		fields = append(fields, n)
	}

	return fields, nil
}

func parseRange(part string) (int, int, error) {
	bounds := strings.SplitN(part, "-", 2)
	if len(bounds) != 2 || bounds[0] == "" || bounds[1] == "" {
		return 0, 0, fmt.Errorf("invalid range: %q", part)
	}

	start, err := strconv.Atoi(bounds[0])
	if err != nil || start <= 0 {
		return 0, 0, fmt.Errorf("invalid range start: %q", bounds[0])
	}

	end, err := strconv.Atoi(bounds[1])
	if err != nil || end <= 0 {
		return 0, 0, fmt.Errorf("invalid range end: %q", bounds[1])
	}

	if start > end {
		return 0, 0, fmt.Errorf("invalid range: %q", part)
	}

	return start, end, nil
}

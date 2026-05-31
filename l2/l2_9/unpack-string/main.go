package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	unpack, err := UnpackString(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(unpack)
}

// UnpackString unpacks a string that contains repeating characters
func UnpackString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(s)

	var prev rune
	escaped := false

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if escaped {
			result.WriteRune(r)
			prev = r
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", errors.New("invalid string")
			}

			count := int(r - '0')

			for j := 0; j < count-1; j++ {
				result.WriteRune(prev)
			}

			prev = 0
			continue
		}

		result.WriteRune(r)
		prev = r
	}

	if escaped {
		return "", errors.New("invalid escape")
	}

	return result.String(), nil
}

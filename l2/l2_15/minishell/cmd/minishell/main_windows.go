//go:build windows

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "minishell: requires Unix (Linux, macOS, or WSL)")
	os.Exit(1)
}

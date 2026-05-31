//go:build unix

package main

import "minishell/internal/shell"

func main() {
	shell.New().Run()
}

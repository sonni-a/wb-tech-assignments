//go:build unix

package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"minishell/internal/parser"
)

func isBuiltin(name string) bool {
	switch name {
	case "cd", "pwd", "echo", "kill", "ps":
		return true
	default:
		return false
	}
}

func runBuiltin(cmd parser.Command, stdin io.Reader, stdout, stderr io.Writer) (int, error) {
	if len(cmd.Args) == 0 {
		return 0, nil
	}

	name := cmd.Args[0]
	args := cmd.Args[1:]

	switch name {
	case "cd":
		return builtinCD(args)
	case "pwd":
		return builtinPWD(stdout)
	case "echo":
		return builtinEcho(args, stdout)
	case "kill":
		return builtinKill(args, stderr)
	case "ps":
		return builtinPS(stdout)
	default:
		return 1, fmt.Errorf("unknown builtin %q", name)
	}
}

func builtinCD(args []string) (int, error) {
	target := ""
	if len(args) == 0 {
		target = os.Getenv("HOME")
		if target == "" {
			fmt.Fprintln(os.Stderr, "cd: HOME not set")
			return 1, nil
		}
	} else {
		target = args[0]
	}
	if err := os.Chdir(target); err != nil {
		fmt.Fprintf(os.Stderr, "cd: %v\n", err)
		return 1, nil
	}
	return 0, nil
}

func builtinPWD(stdout io.Writer) (int, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
		return 1, nil
	}
	fmt.Fprintln(stdout, wd)
	return 0, nil
}

func builtinEcho(args []string, stdout io.Writer) (int, error) {
	fmt.Fprintln(stdout, strings.Join(args, " "))
	return 0, nil
}

func builtinKill(args []string, stderr io.Writer) (int, error) {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "kill: usage: kill <pid>")
		return 1, nil
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(stderr, "kill: invalid pid: %v\n", err)
		return 1, nil
	}
	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
		fmt.Fprintf(stderr, "kill: %v\n", err)
		return 1, nil
	}
	return 0, nil
}

func builtinPS(stdout io.Writer) (int, error) {
	fmt.Fprintf(stdout, "%-8s %s\n", "PID", "COMM")

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return runExternalPS(stdout)
	}

	for _, ent := range entries {
		pid, err := strconv.Atoi(ent.Name())
		if err != nil || pid <= 0 {
			continue
		}
		commPath := fmt.Sprintf("/proc/%d/comm", pid)
		data, err := os.ReadFile(commPath)
		if err != nil {
			continue
		}
		comm := strings.TrimSpace(string(data))
		fmt.Fprintf(stdout, "%-8d %s\n", pid, comm)
	}
	return 0, nil
}

func runExternalPS(stdout io.Writer) (int, error) {
	c := exec.Command("ps", "-eo", "pid,comm")
	c.Stdout = stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return 1, nil
	}
	return 0, nil
}

//go:build unix

package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"minishell/internal/parser"
)

func (sh *Shell) runJob(j *parser.Job) int {
	if j == nil || len(j.Pipelines) == 0 {
		return 0
	}

	runNext := true
	exitCode := 0

	for i := range j.Pipelines {
		if runNext {
			exitCode = sh.runPipeline(&j.Pipelines[i])
		}
		if i >= len(j.Ops) {
			break
		}
		switch j.Ops[i] {
		case "&&":
			runNext = exitCode == 0
		case "||":
			runNext = exitCode != 0
		}
	}
	return exitCode
}

func (sh *Shell) runPipeline(pl *parser.Pipeline) int {
	if len(pl.Commands) == 0 {
		return 0
	}
	if len(pl.Commands) == 1 {
		code, err := sh.runCommand(&pl.Commands[0], os.Stdin, os.Stdout, os.Stderr)
		sh.untrack()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
		return code
	}
	return sh.runPipelineChain(pl.Commands)
}

func (sh *Shell) runPipelineChain(commands []parser.Command) int {
	n := len(commands)
	pipeR := make([]*os.File, n-1)
	pipeW := make([]*os.File, n-1)

	for i := 0; i < n-1; i++ {
		r, w, err := os.Pipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "pipe: %v\n", err)
			return 1
		}
		pipeR[i] = r
		pipeW[i] = w
	}

	type cmdResult struct {
		code int
		err  error
	}
	results := make([]cmdResult, n)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := range commands {
		i := i
		var stdin io.Reader = os.Stdin
		var stdout io.Writer = os.Stdout

		if i > 0 {
			stdin = pipeR[i-1]
		}
		if i < n-1 {
			stdout = pipeW[i]
		}

		go func() {
			defer wg.Done()
			code, err := sh.runCommand(&commands[i], stdin, stdout, os.Stderr)
			if i < n-1 {
				_ = pipeW[i].Close()
			}
			if i > 0 {
				_ = pipeR[i-1].Close()
			}
			results[i] = cmdResult{code: code, err: err}
		}()
	}

	wg.Wait()
	sh.untrack()

	last := results[n-1]
	if last.err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", last.err)
		return 1
	}
	return last.code
}

func (sh *Shell) runCommand(cmd *parser.Command, stdin io.Reader, stdout, stderr io.Writer) (int, error) {
	if len(cmd.Args) == 0 {
		return 0, nil
	}

	in := stdin
	out := stdout

	if cmd.Stdin != nil {
		f, err := os.Open(cmd.Stdin.Path)
		if err != nil {
			return 1, err
		}
		defer f.Close()
		in = f
	}

	if cmd.Stdout != nil {
		flags := os.O_CREATE | os.O_WRONLY
		if cmd.AppendOut {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
		f, err := os.OpenFile(cmd.Stdout.Path, flags, 0o644)
		if err != nil {
			return 1, err
		}
		defer f.Close()
		out = f
	}

	name := cmd.Args[0]
	if isBuiltin(name) {
		return runBuiltin(*cmd, in, out, stderr)
	}

	exe, err := exec.LookPath(name)
	if err != nil {
		return 1, fmt.Errorf("%s: command not found", name)
	}

	c := exec.Command(exe, cmd.Args[1:]...)
	c.Stdin = in
	c.Stdout = out
	c.Stderr = stderr
	c.Env = os.Environ()

	if err := c.Start(); err != nil {
		return 1, err
	}
	sh.track(c.Process)

	err = c.Wait()
	if err == nil {
		return 0, nil
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		}
	}
	return 1, err
}

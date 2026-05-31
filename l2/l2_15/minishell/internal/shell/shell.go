//go:build unix

package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"minishell/internal/parser"
)

const Prompt = "minishell> "

// Shell is a minimal Unix command interpreter.
type Shell struct {
	mu    sync.Mutex
	procs []*os.Process
}

// New creates a shell instance.
func New() *Shell {
	return &Shell{}
}

func (sh *Shell) track(p *os.Process) {
	sh.mu.Lock()
	sh.procs = append(sh.procs, p)
	sh.mu.Unlock()
}

func (sh *Shell) untrack() {
	sh.mu.Lock()
	sh.procs = nil
	sh.mu.Unlock()
}

func (sh *Shell) interruptRunning() {
	sh.mu.Lock()
	defer sh.mu.Unlock()
	for _, p := range sh.procs {
		_ = p.Signal(syscall.SIGINT)
	}
}

// RunLine parses and executes one input line.
func (sh *Shell) RunLine(line string) int {
	j, err := parser.ParseLine(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		return 1
	}
	return sh.runJob(j)
}

// Run starts an interactive loop or reads commands from stdin (script mode).
func (sh *Shell) Run() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	go func() {
		for range sigCh {
			sh.interruptRunning()
		}
	}()

	interactive := isTerminal(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	for {
		if interactive {
			fmt.Print(Prompt)
		}
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if interactive {
				fmt.Println()
			}
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			break
		}

		line = strings.TrimRight(line, "\r\n")
		if strings.TrimSpace(line) == "" {
			continue
		}

		_ = sh.RunLine(line)
	}
}

func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

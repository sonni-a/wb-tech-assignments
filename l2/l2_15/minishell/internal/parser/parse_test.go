package parser

import (
	"os"
	"testing"
)

func TestParsePipeline(t *testing.T) {
	j, err := ParseLine("echo hello | wc -c")
	if err != nil {
		t.Fatal(err)
	}
	if len(j.Pipelines) != 1 || len(j.Pipelines[0].Commands) != 2 {
		t.Fatalf("got %+v", j)
	}
	if j.Pipelines[0].Commands[0].Args[0] != "echo" {
		t.Fatal("expected echo")
	}
}

func TestParseLogical(t *testing.T) {
	j, err := ParseLine("false && echo no || echo yes")
	if err != nil {
		t.Fatal(err)
	}
	if len(j.Pipelines) != 3 || len(j.Ops) != 2 {
		t.Fatalf("got %+v", j)
	}
	if j.Ops[0] != "&&" || j.Ops[1] != "||" {
		t.Fatal("bad ops")
	}
}

func TestParseRedirect(t *testing.T) {
	j, err := ParseLine("echo hi > out.txt")
	if err != nil {
		t.Fatal(err)
	}
	cmd := j.Pipelines[0].Commands[0]
	if cmd.Stdout == nil || cmd.Stdout.Path != "out.txt" {
		t.Fatalf("stdout redirect: %+v", cmd)
	}
}

func TestParseFields(t *testing.T) {
	j, err := ParseLine("echo one two three")
	if err != nil {
		t.Fatal(err)
	}
	args := j.Pipelines[0].Commands[0].Args
	if len(args) != 4 || args[3] != "three" {
		t.Fatalf("strings.Fields args=%v", args)
	}
}

func TestExpandEnv(t *testing.T) {
	_ = os.Setenv("MINISHELL_TEST_VAR", "world")
	defer os.Unsetenv("MINISHELL_TEST_VAR")

	j, err := ParseLine("echo hello $MINISHELL_TEST_VAR")
	if err != nil {
		t.Fatal(err)
	}
	args := j.Pipelines[0].Commands[0].Args
	if len(args) != 3 || args[2] != "world" {
		t.Fatalf("args=%v", args)
	}
}

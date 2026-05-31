#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
go build -o minishell ./cmd/minishell

export MINISHELL_TEST_VAR=expanded

run() {
  printf '%s\n' "$@" | ./minishell
}

echo "== echo =="
out=$(run 'echo hello world')
test "$out" = "hello world"

echo "== pwd =="
run 'pwd' | grep -q .

echo "== cd =="
run 'cd /tmp' 'pwd' | grep -q '^/tmp'

echo "== redirect =="
rm -f /tmp/minishell_out.txt
run 'echo line > /tmp/minishell_out.txt'
grep -q line /tmp/minishell_out.txt

echo "== env =="
out=$(run 'echo $MINISHELL_TEST_VAR')
test "$out" = "expanded"

echo "== pipeline =="
out=$(run 'echo -e "a\nb\nc" | wc -l')
test "$out" -ge 1

echo "== && / || =="
out=$(run 'false && echo no' 'echo after-fail')
test "$out" = "after-fail"
out=$(run 'true || echo no' 'echo after-true')
test "$out" = "after-true"

echo "== external =="
run 'ls' >/dev/null

echo "All tests passed."

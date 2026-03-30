package main

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	duration := 100 * time.Millisecond

	start := time.Now()
	Sleep(duration)
	elapsed := time.Since(start)

	if elapsed < duration {
		t.Errorf("got %v, want at least %v", elapsed, duration)
	}
}

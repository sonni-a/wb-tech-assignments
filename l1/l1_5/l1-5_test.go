package main

import (
	"testing"
	"time"
)

func TestTimeoutChannel(t *testing.T) {
	duration := 1 * time.Second
	ch := timeoutChannel(duration)

	start := time.Now()

	count := 0
	for range ch {
		count++
	}

	elapsed := time.Since(start)

	if elapsed < duration {
		t.Errorf("channel closed too soon: %v", elapsed)
	}

	if count == 0 {
		t.Error("didn't get anything")
	}
}

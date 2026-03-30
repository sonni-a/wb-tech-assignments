package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestWorkersCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	worked := make([]bool, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					worked[id] = true
					time.Sleep(10 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()

	for i, w := range worked {
		if !w {
			t.Errorf("Worker %d did not work", i)
		}
	}
}

package main

import (
	"sync"
	"testing"
)

func TestSafeMapConcurrent(t *testing.T) {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			sm.Set(i, i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 1000; i++ {
		val, ok := sm.Get(i)
		if !ok || val != i {
			t.Errorf("wrong value for key %d", i)
		}
	}
}

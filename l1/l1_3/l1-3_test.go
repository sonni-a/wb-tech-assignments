package main

import (
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	jobs := make(chan Job)
	results := make(chan int)

	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job.ID
			}
		}()
	}

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- Job{ID: i}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for id := range results {
		count++
		t.Logf("Processed Job ID: %d", id)
	}

	if count != 5 {
		t.Errorf("expected %d jobs processed, got %d", 5, count)
	}
}

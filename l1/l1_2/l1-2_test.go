package main

import (
	"sync"
	"testing"
)

func TestSquares(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	res := make(chan int, len(nums))

	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			res <- num * num
		}(n)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	got := []int{}
	for r := range res {
		got = append(got, r)
	}

	if len(got) != len(nums) {
		t.Fatalf("got %d results, want %d", len(got), len(nums))
	}
}

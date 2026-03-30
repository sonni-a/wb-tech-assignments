package main

import (
	"reflect"
	"testing"
)

func readAll(ch <-chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestPipeline(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	ch1 := generateNumbers(input)
	ch2 := proccessNumbers(ch1)

	result := readAll(ch2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

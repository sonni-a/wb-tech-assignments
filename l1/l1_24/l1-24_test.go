package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	result := p1.Distance(p2)
	expected := 5.0

	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

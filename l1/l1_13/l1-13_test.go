package main

import "testing"

func TestSwapXOR(t *testing.T) {
	a, b := 5, 10
	x, y := swapXOR(a, b)

	if x != b || y != a {
		t.Errorf("swapXOR failed: got (%d, %d), want (%d, %d)", x, y, b, a)
	}
}

func TestSwapSum(t *testing.T) {
	a, b := 5, 10
	x, y := swapSum(a, b)

	if x != b || y != a {
		t.Errorf("swapSum failed: got (%d, %d), want (%d, %d)", x, y, b, a)
	}
}

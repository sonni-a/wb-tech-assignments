package main

import (
	"math/big"
	"testing"
)

func TestOperations(t *testing.T) {
	a := big.NewInt(1 << 25)
	b := big.NewInt(1 << 22)

	t.Run("Add", func(t *testing.T) {
		expected := new(big.Int).Add(a, b)
		result := Add(a, b)

		if result.Cmp(expected) != 0 {
			t.Errorf("Add failed: got %v, want %v", result, expected)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		expected := new(big.Int).Sub(a, b)
		result := Sub(a, b)

		if result.Cmp(expected) != 0 {
			t.Errorf("Sub failed: got %v, want %v", result, expected)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		expected := new(big.Int).Mul(a, b)
		result := Mul(a, b)

		if result.Cmp(expected) != 0 {
			t.Errorf("Mul failed: got %v, want %v", result, expected)
		}
	})

	t.Run("Div", func(t *testing.T) {
		expected := new(big.Int).Div(a, b)
		result := Div(a, b)

		if result == nil || result.Cmp(expected) != 0 {
			t.Errorf("Div failed: got %v, want %v", result, expected)
		}
	})
}

func TestDivisionByZero(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(0)

	result := Div(a, b)
	if result != nil {
		t.Errorf("Expected nil for division by zero, got %v", result)
	}
}

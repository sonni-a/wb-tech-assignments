package main

import "testing"

func TestTypeRuntime(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"int", 10, "int"},
		{"string", "hello", "string"},
		{"bool", true, "bool"},
		{"chan int", make(chan int), "chan"},
		{"chan string", make(chan string), "chan"},
		{"unknown", 3.14, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := typeRuntime(tt.input)
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

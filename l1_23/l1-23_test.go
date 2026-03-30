package main

import (
	"reflect"
	"testing"
)

func TestDeleteElement_Int(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := deleteElement(input, 1)

	expected := []int{1, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestDeleteElement_String(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := deleteElement(input, 0)

	expected := []string{"b", "c"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

type User struct {
	Name string
}

func TestDeleteElement_Struct(t *testing.T) {
	input := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}

	result := deleteElement(input, 1)

	expected := []User{
		{Name: "Alice"},
		{Name: "Charlie"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

package main

import "testing"

func TestGetOlder(t *testing.T) {
	h := Human{Name: "Sonya", Age: 20}

	h.GetOlder()

	if h.Age != 21 {
		t.Errorf("expected %d, got %d", 21, h.Age)
	}
}

func TestGetSomething(t *testing.T) {
	h := Human{Name: "Sonya"}

	h.GetSomething("labubu")

	if len(h.Items) != 1 || h.Items[0] != "labubu" {
		t.Errorf("expected %v, got %v", "labubu", h.Items)
	}
}

func TestActionUsesHumanMethods(t *testing.T) {
	a := Action{
		Human: Human{Name: "Sonya", Age: 20},
		Event: "Birthday",
	}

	a.GetOlder()
	if a.Age != 21 {
		t.Errorf("expected %d, got %d", 21, a.Age)
	}

	a.GetSomething("labubu")
	if len(a.Items) != 1 || a.Items[0] != "labubu" {
		t.Errorf("expected %v, got %v", "labubu", a.Items)
	}
}

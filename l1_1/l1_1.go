package main

import (
	"fmt"
)

type Human struct {
	Name  string
	Age   int
	Items []string
}

func (h *Human) GetOlder() {
	h.Age++
	fmt.Printf("%s is %d years old\n", h.Name, h.Age)
}

func (h *Human) GetSomething(item string) {
	h.Items = append(h.Items, item)
	fmt.Printf("Woah, %s got a %s!\n", h.Name, item)
}

type Action struct {
	Event string
	Human
}

func main() {
	a := Action{
		Human: Human{Name: "Sonya", Age: 20},
		Event: "Birthday",
	}

	fmt.Println("Activity:", a.Event)
	a.GetOlder()
	a.GetSomething("labubu")
}

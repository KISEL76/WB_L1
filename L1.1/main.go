package main

import "fmt"

type Human struct {
	height int
	age    int
	name   string
}

type Action struct {
	Human
}

func (h *Human) SayAge() {
	fmt.Printf("Hello, I am %d y.o.\n", h.age)
}

func (h *Human) GetAcquainted() {
	fmt.Printf("Hello, I am %s and I am %d years old. And my height is %d\n", h.name, h.age, h.height)
}

func NewAction(age, height int, name string) *Action {
	return &Action{
		Human{
			height: height,
			age:    age,
			name:   name,
		},
	}
}

func main() {
	Johny := NewAction(21, 180, "Johny")
	Johny.SayAge()
	Johny.GetAcquainted()
}

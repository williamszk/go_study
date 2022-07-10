package main

import "fmt"

type Animal struct {
	Name string
	Size int
}

func (a Animal) HandleFromInside() {
	fmt.Printf("Calling from inside HandleFromInside(): \t\t%p\n", &a)
}

// In C:
// Animal *a
// int *ptr
func (ptr *Animal) HandleFromInsidePtrReceiver() {
	// this first one is about the pointer itself, not the object that it points to
	fmt.Printf("Calling from inside HandleFromInsidePtrReceiver() ptr: \t%p\n", &ptr)
	fmt.Printf("Calling from inside HandleFromInsidePtrReceiver(): \t%p\n", ptr)
}

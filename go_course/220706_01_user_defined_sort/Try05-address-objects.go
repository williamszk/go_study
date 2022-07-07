// Experiment with observing the address of objects
// when we don't use pointer receiver
// for: int, struct and slice
package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) MakeYoung2() {
	p.Age = 10
	fmt.Println(p.Name, "is", p.Age, "years old.")
	fmt.Printf("Inside MakeYoung2 func, p address: 0x%p\n", &p)
}

func Try05() {

	person := Person2{
		Name: "Bob",
		Age:  90,
	}

	fmt.Printf("person address: 0x%p\n", &person)
	person.MakeYoung2()
	fmt.Printf("person.Name address: 0x%p\n", &person.Name)
}

// person address: 0x&{426F62 5A}
// Bob is 10 years old.
// Inside MakeYoung2 func, p address: 0x&{426F62 A}
// person.Name address: 0xC000004078

// Why the address of a struct is in a difference format compared
// to the address of element datatypes

// We need to use %p to print the address
// person address: 0x0xc000004078
// Bob is 10 years old.
// Inside MakeYoung2 func, p address: 0x0xc000004090
// person.Name address: 0x0xc000004078

// Note that the address of the struct is the same as the address
// of its first element

// experiment with checking if structs when used as an argument to a function
// it is passed as a reference or the whole struct is copied.

// what happens if the struct have a slice as a field?
package main

import "fmt"

func main() {

	bob := person{
		name:     "Bob",
		age:      10,
		iceCream: []string{"Chocolate", "Strawberry"},
	}

	fmt.Println("Before: ", bob)
	modifyPerson(bob)
	fmt.Println("After: ", bob)

	// Before:  {Bob 10 [Chocolate Strawberry]}
	// After:  {Bob 10 [Chocolate Strawberry]}
}

type person struct {
	name     string
	age      int
	iceCream []string
}

func modifyPerson(p person) {
	// internally modify the struct
	p.name = "modified name"
	p.age = 100
	p.iceCream = []string{"banana"}
	fmt.Println("Internally:", p)
	// Before:  {Bob 10 [Chocolate Strawberry]}
	// Internally: {modified name 100 [banana]}
	// After:  {Bob 10 [Chocolate Strawberry]}
}

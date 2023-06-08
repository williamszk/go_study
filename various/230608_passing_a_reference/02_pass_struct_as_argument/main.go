// experiment with checking if structs when used as an argument to a function
// it is passed as a reference or the whole struct is copied.
package main

import "fmt"

func main() {

	bob := person{
		name: "Bob",
		age:  10,
	}

	fmt.Println("Before: ", bob)
	modifyPerson(bob)
	fmt.Println("After: ", bob)
	// Before:  {Bob 10}
	// After:  {Bob 10}

	// passing value will not modify the struct
	// we would need to pass the reference to a struct
}

type person struct {
	name string
	age  int
}

func modifyPerson(p person) {
	// internally modify the struct
	p.name = "modified name"
	p.age = 100
}

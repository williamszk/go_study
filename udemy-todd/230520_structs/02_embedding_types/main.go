package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// `secretAgent` has all the fields of a `person` but it also has licenseToKill
// if we change `person` we change `secretAgent`
type secretAgent struct {
	person
	licenseToKill bool
}

type AbleToKill interface {
	kill(target person) (ok bool)
}

func (s secretAgent) kill(target person) (ok bool) {
	fmt.Println(s.first, s.last, "killed", target.first, target.last)
	return true
}

type assassin struct {
	AbleToKill
}

func main() {

	// when we instantiate a value of `secretAgent` we need to explicitly declare
	// the person field.
	//
	jamesBond := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		licenseToKill: true,
	}
	// but when we reach for the secretAgent's fields we can call them directly
	fmt.Println(jamesBond.first, jamesBond.last, jamesBond.age)

	// -------------------------------------------------------------------------

	// About embedding structs and promoting fields:
	// when we embed structs inside structs it is like a inheritance, but a little different
	// when creating the value, it is as if the embedded struct (or the inner type) is a just another field,
	// which by happenstance is a struct.
	// When calling the fields of the inner type, we say that the fields of the embedded structs are promoted,
	// so the whole thing behaves like the inner type fields belong to the outer type.

	// -------------------------------------------------------------------------

	// What is the difference between:
	// 1. Embedding struct inside struct.
	// 2. Embedding interface inside struct.
	// 3. Embedding interface inside another interface.
	//
	// For case (1) is what we already read above.
	// For case (2) it is a more general way of the (1), in the sense that
	// instead of embedding just a single concrete type, we can create structs which embed any type
	// that implements a certain interface.
	// Question: Are the methods (functions) of the embedded interface (inner type) promoted to the
	// embedding struct (outer type)?
	// Yes, after experimentation, the methods from the embedded interface (inner type)
	// are promoted to the embedding struct (outer type).
	// I believe this also happens with the case (3). In which the methods of the
	// embedded interface (inner type) are promoted to the embedding interface (outer type).

	func() string {
		return "Hi!"
	}()

	bobTheAssassin := assassin{
		AbleToKill: secretAgent{
			person: person{
				first: "Bob",
				last:  "Ferguson",
				age:   88,
			},
			licenseToKill: false,
		},
	}

	fmt.Println(bobTheAssassin)

	// bobTheAssassin.kill(jamesBond) // this will give an error
	// Bob Ferguson can't kill James Bond because James Bond is a `secretAgent` and not a `person`...
	// Which if we are not thinking strictly about types, it doesn't make sense.
	// Ideally person should be an interface.

	garyFields := person{
		first: "Gary",
		last:  "Fields",
		age:   22,
	}

	bobTheAssassin.kill(garyFields)
	// We can see here the the methods from the `ableToKill` interface were promoted
	// to the `assassin` struct.

	// -------------------------------------------------------------------------
}

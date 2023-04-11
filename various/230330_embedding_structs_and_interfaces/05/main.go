package main

import "fmt"

func main() {

	a := Astruct{
		Bstruct: Bstruct{
			B1: "This is from B1 inside Astruct",
		},
		A1: "This is field A1 from Astruct",
	}

	// foo(a) // we can't achieve polymorphism using struct inside struct
	// but maybe we can achieve using interface inside struct

	bar(a)
	// bar receives an Interface Binterface.
	// Astruct has embedded a Bstruct which implements the Binterface
	// An alternative to achieve polymorphism is to use an interface embedded
	// inside a struct.
	// For example we could embed Binterface inside the Astruct and this would be
	// a more flexible alternative where the field inside Astruct is not specifically
	// a Bstruct, but we could implement another Bstruct2 which implements the
	// Binterface and use it inside the Astruct.

	// Alway we can use interfaces instead of concrete structs where type are
	// being used to either set the parameters or the return value.

	a2 := Astruct2{
		Binterface: Bstruct{
			B1: "This is a field inside the Bstruct inside the Astruct2",
		},
		A2: "Hi from Astruct2",
	}

	bar(a2)

}

type Astruct struct {
	Bstruct
	A1 string
}

type Bstruct struct {
	B1 string
}

func (b Bstruct) Bmethod() {
	fmt.Println("from Bstruct Bmethod: " + b.B1)
}

type Binterface interface {
	Bmethod()
}

func foo(b Bstruct) {
	fmt.Println(b.B1)
}

func bar(b Binterface) {
	// we are including inside the Binterface as the type receiver
	// but we can pass any struct that implements the Binterface
	// or we can pass a struct which embedded has the Binterface or
	// a struct which has embedded another struct which implements the
	// Binterface.
	fmt.Println("This is being called from bar")
}

type Astruct2 struct {
	Binterface
	A2 string
}

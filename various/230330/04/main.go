package main

import "fmt"

func main() {

	a := Astruct{
		Bstruct: Bstruct{
			B1: "This is from B1 inside Astruct",
		},
		A1: "This is field A1 from Astruct",
	}

	foo(a) // we can't achieve polymorphism using struct inside struct
	// but maybe we can achieve using interface inside struct

}

type Astruct struct {
	Bstruct
	A1 string
}

type Bstruct struct {
	B1 string
}

func foo(b Bstruct) {
	fmt.Println(b.B1)
}

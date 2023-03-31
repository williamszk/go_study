package main

import "fmt"

func main() {
	b := B{
		b:  "Hi there",
		b1: 99,
	}
	fmt.Println(b)
	fmt.Println(b.concat())

	//
	a := A{
		B: B{
			b:  "Hi from B inside A",
			b1: 1888,
		},
		a: "This is a string from A inside A",
	}

	fmt.Println(a)

	fmt.Println(a.concat())

}

type A struct {
	B // <--- embed struct inside struct, this will work like inheritance, of a specific kind
	a string
}

type B struct {
	b  string
	b1 int
}

func (b B) concat() string {
	return fmt.Sprintf("%v %v", b.b, fmt.Sprintf("%v", b.b1))
}

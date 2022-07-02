// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922290#overview

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
	// with this pattern the code will know that Goperh and Zebra
	// are aslo of Animal type and the values that they receive are next to
	// iota, which can be 1, 2, 3 etc
)

// UnmarshalJSON is a method of  Animal the * at the Animal
// doesn't mean that "a" is a pointer to Animal, it is just to Animal
// "b" is a slice of bytes, for what "b" is used for?
// and the return of the function is just error
func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	// here we run two lines of code in one
	// if err := json.Unmarshal(b, &s); err != nil {
	// 	return err
	// }

	// we need to pass the address of "s" to the Unmarshal function
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// "strings" is the package, I imagine
	switch strings.ToLower(s) {
	// we are changing the value of the object
	// Animal, which is an int
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func main() {

	fmt.Println("Unknown: ", Unknown)
	fmt.Println("Gopher: ", Gopher)
	fmt.Println("Zebra: ", Zebra)

	// fmt.Printf("%v\t%T\n", error, error)
	// error is a builtin interface

	a := Unknown

	fmt.Println("before a: ", a)

	myAnimal := "gopher"

	// how to convert string into a slice of bytes in Go?
	// how to do this in Python?
	b := []byte(myAnimal)

	fmt.Println("b: ", b)

	err := a.UnmarshalJSON(b)
	// I got some error here that I couldn't understand
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after a: ", a)
}

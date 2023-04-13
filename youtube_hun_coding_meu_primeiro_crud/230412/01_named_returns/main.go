package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hi there!")

	name, id := AnExampleOfNamedReturn()

	fmt.Printf("name: %v, id: %v\n", name, id)

	out, err := ReturnValueErr()

	if err != nil {
		log.Printf("We got an error: %v\n", err)
	}

	fmt.Println("out:", out)

}

// The named returns will declare the values right at the start of
// the function, so we just assign the values directly in the body
func AnExampleOfNamedReturn() (name string, id int) {

	name = "Alice"
	id = 82713

	return name, id
}

func ReturnValueErr() (out int, err error) {
	a := 10
	if a < 12 {
		err = errors.New("invalid value")
		return out, err
	} else {
		out = 10
		err = nil
		return out, err
	}
}

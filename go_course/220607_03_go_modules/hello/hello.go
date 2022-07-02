package main

import (
	"fmt"

	"example.com/greetings"
)

// we need to run this command so go can find the module
// go mod edit -replace example.com/greetings=../greetings
// in production we should publish example.com/greetings
// we need to run
// go mod tidy
func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

// we can use to run inside the package main
// go run .
//

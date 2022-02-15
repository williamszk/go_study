package main

import "fmt"

func main() {
	defer fmt.Println("printing 1")
	defer fmt.Println("printing 2")
	fmt.Println("printing 3")
	fmt.Println("printing 4")

	// printing 3
	// printing 4
	// printing 2
	// printing 1

	// defer is used to close a file after all the code has run
	// open(file)
	// defer close(file)
	// it is easier to remember to close the file right after we open it
}

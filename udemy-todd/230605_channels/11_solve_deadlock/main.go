package main

import "fmt"

// solve the deadlock
func main() {

	// we could also have used buffered channels
	// c := make(chan int, 1)
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c) // this will block the execution of the program

}

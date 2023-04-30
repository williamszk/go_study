package main

import "fmt"

// make this function work
// 1. using an anonymous function
// 2. using buffers

// here we'll use buffers to solve

func main() {
	c := make(chan int, 1)
	// In Go, channels are used for communication between concurrent goroutines. 
	// Buffers can be added to channels to allow them to store a limited number of 
	// values. This can be useful in situations where the sending or receiving of 
	// values on a channel may be slower than the other end.

	c <- 24

	fmt.Println(<-c)
}

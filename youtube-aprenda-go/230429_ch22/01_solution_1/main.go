package main

import "fmt"

// make this function work
// 1. using an anonymous function
// 2. using buffers

// We'll use anonymous functions here
func main() {
	c := make(chan int)

	go func(){
		c <- 24
	}()

	fmt.Println(<-c)
}

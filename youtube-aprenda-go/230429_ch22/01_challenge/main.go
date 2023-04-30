package main

import "fmt"

// make this function work
// 1. using an anonymous function
// 2. using buffers

func main() {
	c := make(chan int)

	c <- 24

	fmt.Println(<-c)
}

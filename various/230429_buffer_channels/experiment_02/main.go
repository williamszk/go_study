package main

import "fmt"

// What happens when we create a buffered channel and no one sends data into it?
func main() {

	myChan := make(chan int, 2)

	myVal := <-myChan

	fmt.Println(myVal) // <- this will give a deadlock error
	// this happens because main is waiting for some value to enter the channel
	// but it knows that no other go routine is running that could send values
	// into it.

}

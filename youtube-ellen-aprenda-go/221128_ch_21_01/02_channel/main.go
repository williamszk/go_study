package main

import "fmt"

// directed channels
func main() {
	// a bi-directed channel
	myChan := make(chan int)

	go send(myChan)

	receive(myChan)

}

// s is a channel
// and s needs to be a channel that will only receive ints
func send(s chan<- int) {
	s <- 42
}

// r is a channel, that can only receive
// values of type int
func receive(r <-chan int) {
	fmt.Println("The value received from the channel is:", <-r)
}

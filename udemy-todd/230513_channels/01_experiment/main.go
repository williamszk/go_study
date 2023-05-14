// Experiment with buffered and un-buffered channels
// with an un-buffered channel, the channel can hold only 1 item
// if no one

package main

import (
	"fmt"
	"time"
)

func main() {

	// this is an un-buffered channel
	c := make(chan int)

	go func() {
		fmt.Println("Printing from a go routine")
		for i := 0; i < 10; i++ {
			fmt.Println("A chan<-", i)
			c <- i // <-- the program will stop here, if no one takes out the value from the channel
		}
	}()
	// Printing from a go routine
	// chan<- 0
	// it stops here

	// ========================================================================
	fmt.Println("------------------------------------------------")

	// build a buffered channel
	c2 := make(chan int, 3)
	go func() {
		fmt.Println("Printing from a go routine 2")
		for i := 0; i < 10; i++ {
			fmt.Println("B chan<-", i)
			c2 <- i
		}
	}()
	// Printing from a go routine 2
	// B chan<- 0
	// B chan<- 1
	// B chan<- 2
	// B chan<- 3

	// this channel can hold only 4 items

	time.Sleep(10 * time.Second)

	// Summary
	// by default channels (the un-buffered channel) can hold only
	// 1 item, and the goroutine will halt at the line of input
	// until another goroutine consumes its value

	// that is the un-buffered channels block
	// the operation <-c or c<- will block the execution
	// of the go routine until something is fed or taken from it
}

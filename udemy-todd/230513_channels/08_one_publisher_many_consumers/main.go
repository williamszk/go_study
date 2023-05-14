// In here we build one producer and two consumers

package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100_000; i++ {
			c <- i // remember that the un-buffered channel can only hold one item
			// only after this item has been consumed, the producer can insert
			// another item
		}

		// after the producer has finished inserting into the channel
		// then it can close the channel
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done // the program will halt here and wait for the consumers to send the
	<-done // signal that they finished consuming. this will hold the main goroutine
}

// this is a program that creates two consumers to consume from one
// producer... we could build a program that compares the speed when
// we have more consumers... we could use the benchmark functionality
// from the go standard library

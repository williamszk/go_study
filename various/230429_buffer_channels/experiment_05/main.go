package main

import (
	"fmt"
	"time"
)

// What happens when we create a buffered channel and the sender is taking too
// long to send data?
// Somehow Go knows that some other goroutine is going to send some data to the
// channel...
// Or if the other goroutine is sending to another channel?
func main() {

	myChan := make(chan int)
	myChan2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second) // wait for 3 seconds
		myChan2 <- 10               // note that we're sending to another channel
	}()

	myVal := <-myChan
	// go doesn't know if the goroutine running is related or not to this read
	// it will wait to see if any other go-routine is running
	// if no more go routines are running
	// and it tries to read this value, and no value is in the channel
	// the program will panic

	// What if then we have lots of go routines being create indefinitely?
	fmt.Println(myVal)
}

package main

import (
	"fmt"
	"time"
)

// What happens when we create a buffered channel and the sender is taking too
// long to send data?
func main() {

	myChan := make(chan int)

	go func() {
		time.Sleep(3 * time.Second) // wait for 3 seconds
		myChan <- 10
	}()

	myVal := <-myChan

	fmt.Println(myVal)
	// Somehow Go knows that some other goroutine is going to send some data to the
	// channel...
	// What if there is a chance that this happens?
	// Or if the other goroutine is sending to another channel?
}

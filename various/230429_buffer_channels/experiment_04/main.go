package main

import (
	"fmt"
	"math/rand"
	"time"
)

// What happens when we create a buffered channel and the sender is taking too
// long to send data?
// Somehow Go knows that some other goroutine is going to send some data to the
// channel...
// What if there is a chance that this happens?
func main() {

	myChan := make(chan int)

	go func() {
		rand.Seed(time.Now().UnixNano()) // set seed for random number generation

		randomInt := rand.Intn(101) // generate a random value between 0 and 100

		time.Sleep(3 * time.Second) // wait for 3 seconds

		// there is a chance that the channel will be inserted with some value
		// but also there is a change that it will not be filled
		if randomInt >= 50 {
			myChan <- 10
		}

	}()

	myVal := <-myChan
	// the result is that go will wait and see
	// if there are no values to be read and no other goroutine is running
	// then the program will panic

	fmt.Println(myVal)
}

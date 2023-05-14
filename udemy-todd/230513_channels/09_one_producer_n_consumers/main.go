package main

import "fmt"

// this program have one producer, and it has many consumers

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100_000; i++ {
			c <- i // this line will halt the goroutine until a consumer picks the
			// item from the channel
		}
		close(c)
	}()

	// create many goroutine consumers
	for j := 0; j < n; j++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	// wait for the consumers to finish
	for j := 0; j < n; j++ {
		<-done // this flag will halt the main goroutine until all the consumer
		// goroutines signal that they have finished consuming
		// then we can finish the program
	}

}

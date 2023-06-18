package main

import "fmt"

func main() {
	// build a goroutine that inputs numbers into a channel
	// build another goroutine that pulls items of the channel

	c := make(chan int)
	done := make(chan bool) // this is a semaphore

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for v := range c {
			fmt.Println("Consume: ", v)
		}
		done <- true
		close(done)
	}()

	for v := range done {
		fmt.Println("done: ", v)
	}
}

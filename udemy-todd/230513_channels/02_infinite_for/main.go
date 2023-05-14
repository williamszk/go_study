package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 7; i++ {
			c <- i
		}
	}()

	go func() {
		// this is an infinite loop
		for {
			fmt.Println(<-c)
		}
		fmt.Println("end")
	}()
	// it is not totally clear for me why we would do this
	// the problem with this code is that the for will never end
	// instead we should use close() and a range clause for the channel

	// this sleep is needed so that we can
	time.Sleep(time.Second * 3)
}

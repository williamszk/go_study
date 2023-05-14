package main

import "fmt"

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool) // this is the semaphore

	// create n number of publisher goroutines
	for i := 1; i <= n; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- j + i*100
			}
			done <- true // this is the signal to tell the program that the goroutine
			// has done its job
		}(i)
	}

	go func() {
		// this for loop will make sure we have finished all goroutines
		// that are publishing into the channel, before we can close the channel
		for i := 0; i < n; i++ {
			<-done // this will free up the space in the channel so that the next
			// publisher goroutine can publish in the channel c
		}
		close(c)
	}()

	// this part can't be a goroutine because we need keep the main goroutine
	// running, the range clause will wait for the close(c) before it finishes
	for n := range c {
		fmt.Println(n)
	}
}

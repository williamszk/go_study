// This program will create some values in two different goroutines
// Then we'll use the fan-in pattern to gather them into one single
// goroutine.
// then we'll monitor for a close and range to use all the produced values
package main

import (
	"fmt"
	"sync"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println("Output from fanIn", v)
	}
	fmt.Println("About to exit")
}

// fanIn here needs to be a send channel because we'll need to close it
// closing a channel is like sending information
// we can't close a channel if it is only receive channel
// but note that 'e' and 'o' are receive channels
func receive(e, o <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}

func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(o)
	close(e)
}

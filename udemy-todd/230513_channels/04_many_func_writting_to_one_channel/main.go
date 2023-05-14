package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	// TODO: are wait groups like mutex? or locks?
	// I think they are not like locks.
	// But actually they will lock the input into a channel if it is full.
	// So maybe they work like locks.

	// create two go-routines that write to the same channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i + 100
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i + 200
		}
		wg.Done()
	}()

	// create a go-routine responsible for closing the channel
	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}

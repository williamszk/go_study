package main

import (
	"fmt"
	"sync"
)

func main() {

	evenChan := make(chan int)

	oddChan := make(chan int)

	converges := make(chan int)

	go send(evenChan, oddChan)

	go receive(evenChan, oddChan, converges)

	for v := range converges {
		fmt.Println("Received value:", v)
	}

}

func send(evenChan, oddChan chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			evenChan <- n
		} else {
			oddChan <- n
		}
	}

	close(evenChan)
	close(oddChan)
}

func receive(evenChan, oddChan, converges chan int) {

	var waitgroups sync.WaitGroup
	waitgroups.Add(2)

	go func() {
		for v := range evenChan {
			converges <- v
		}
		waitgroups.Done()
	}()

	go func() {
		for v := range oddChan {
			converges <- v
		}
		waitgroups.Done()
	}()

	waitgroups.Wait()
	close(converges)

}

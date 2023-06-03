package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	// we want to run a function asynchronously and then wait for its response
	ch := make(chan int)
	go func() {
		// this is a go routine which will fire and doSomething
		ch <- doSomething(5)
		// when doSomething is ready it will feed the channel with the precious
		// information, later this information will be passed to another goroutine
	}()

	theValue := <-ch // the execution of the program will stop here, until theValue
	// is available to be consumed
	fmt.Println(theValue)
}

package main

import "fmt"

func main() {

	// create a channel
	myChan := make(chan int)
	// we need to specify what are the types of information that we
	// need to include inside the channel

	// to include an information inside a channel
	// we need a go routine
	go func() {
		// include information inside channel
		myChan <- 10
	}()

	// To get out a value from a channel
	fmt.Println(<-myChan)
}

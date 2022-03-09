package main

import "fmt"

func main() {
	chn1 := make(chan int)
	go func() {
		chn1 <- 42
	}()

	fmt.Println(<-chn1)

	// --------------------
	// idk why this works
	chn2 := make(chan int, 1) // we are using buffer here
	chn2 <- 43
	// chn2 <- 43 // this will cause deadlock error
	fmt.Println(<-chn2)

}

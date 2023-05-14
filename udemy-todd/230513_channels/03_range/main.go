package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c) // without this close, the 'end' would never be printed
		// the close signals the range that it should end the loop
	}()

	// the range will block the execution
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("end")

	// this sleep is not needed actually
	time.Sleep(time.Second * 10)
}

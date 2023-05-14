package main

import "fmt"

func main() {
	// we can use the done channel to signal the closing of the channels
	// instead of using wait groups

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i + 100
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i + 200
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
		close(channel)
	}()

	v, ok := <-channel

	fmt.Println(v, ok)

	v, ok = <-channel
	fmt.Println(v, ok)

}

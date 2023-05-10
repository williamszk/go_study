package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 10
	}()

	v, ok := <-c
	fmt.Println(v, ok) // 10 true

	go func() {
		c <- 20
	}()

	close(c)

	v, ok = <-c
	fmt.Println(v, ok) // 0 false

}

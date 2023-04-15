package main

import "fmt"

func main() {

	// this creates a reference to a channel
	// so ch is just a pointer with some appended info
	ch := make(chan int)
	// this is a send and receive channel

	fmt.Println(ch)
	go func() {
		ch <- 10
	}()

	fmt.Println(<-ch)

	v := struct{}{}
	// struct{} <---- this is a type
	// struct{}{} <--- this is an expression, which returns an empty struct {}
	fmt.Printf("%v\n", v)

}

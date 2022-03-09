package main

import "fmt"

func main() {
	// send channel
	// receive channel
	// the default is the bidirectional channel

	channel01 := make(chan int) // this is a bidirectional channel

	go send(channel01) // this func needs to be a go routine, idk why

	receive(channel01)

}

func send(sendChan chan<- int) {
	// inject int in the channel
	sendChan <- 42 // inject value 42 in the channel
}

func receive(reciChan <-chan int) {
	// extract int values from this channel
	fmt.Println("The channel value is:", <-reciChan)
}

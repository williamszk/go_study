package main

import "fmt"

func main() {

	myChan := make(chan int)

	go populateChannel(10, myChan)

	getChannelContent(myChan)

}

func populateChannel(total int, snd chan<- int) {
	for i := 0; i < total; i++ {
		snd <- i
	}
	close(snd)
}

func getChannelContent(rcv <-chan int) {
	for v := range rcv {
		fmt.Println(v)
	}
}

package main

import "fmt"

func main() {
	chn := make(chan int)

	go myloop(10, chn)

	custprint(chn)

}

func myloop(total int, chnSend chan<- int) {
	for i := 0; i < total; i++ {
		chnSend <- i
	}
	close(chnSend) // this is necessary to avoid the error of the channel being waiting for more information
}

func custprint(chnRecv <-chan int) {
	for v := range chnRecv { // we include the channel as as iterable object
		fmt.Println("Received from channel:", v)
	}
}

// next:
// https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=149&ab_channel=AprendaGo

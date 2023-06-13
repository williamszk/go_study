package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	// while the send goroutine is sending data to the channels
	// the receive channel will block the main goroutine waiting for data to
	// arrive in the respective channels
	// the quit channel works as a semaphore telling the program to finish
	// note that we didn't close the channels...
	// when the main goroutine finishes, all the other goroutines will be terminated

	receive(even, odd, quit)

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Receiving from even channel\t", v)
		case v := <-o:
			fmt.Println("Receiving from odd channel\t", v)
		case v := <-q:
			fmt.Println("Receiving from quit channel\t", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
}
